package cli

import (
	"context"
	"fmt"
	commobj "github.com/qingcc/yi/commobj/utils"
	"github.com/smallnest/rpcx/client"
	"log"
	"sync"
	"time"
)

type TransServiceClient struct {
	cli client.XClient
}

var (
	transServiceClient     *TransServiceClient
	transServiceClientOnce sync.Once
)

func GetTransServiceClient() *TransServiceClient {
	transServiceClientOnce.Do(func() {
		transServiceClient = new(TransServiceClient)
		servicePath := "transfer"
		var addr = []string{"47.112.210.86:7379"}
		d := client.NewEtcdDiscovery("ali", servicePath, addr, nil)
		transServiceClient.cli = client.NewXClient(servicePath, client.Failtry, client.WeightedRoundRobin, d, client.DefaultOption)
	})
	return transServiceClient
}

func (c *TransServiceClient) doRequest(method string, args interface{}, reply interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := c.cli.Call(ctx, method, args, &reply); err != nil {
		log.Println(fmt.Sprintf("[error] failed to call %s, err:%s", method, err))
	}

	return nil
}

func (c *TransServiceClient) Trans(args commobj.Args) (reply commobj.Reply) {
	err := c.doRequest("Trans", args, reply)
	CheckRpcError(err, &reply.BaseServiceResult)
	return
}

func CheckRpcError(err error, res *commobj.BaseServiceResult) {
	if err != nil {
		res.Success = false
		res.Message = "rpcx call failure:" + err.Error()
		log.Println("[error] rpcx call failure:", err)
	}
}
