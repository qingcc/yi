package logserviceutil

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	addr = "tcp@47.112.210.86:6012"
	auth = "123456"
)

//todo 模拟的日志写入rpc， 未完成
func PushLog2Db(ctx context.Context, logData interface{}) {
	d := client.NewPeer2PeerDiscovery(addr, "")
	op := client.DefaultOption
	xclient := client.NewXClient("logService", client.Failtry, client.RandomSelect, d, op)
	defer xclient.Close()

	xclient.Auth(auth)
	if err := xclient.Call(ctx, "PushOrderLog", logData, nil); err != nil {
		log.Println("日志rpc调用失败：", err.Error())
	}
}
