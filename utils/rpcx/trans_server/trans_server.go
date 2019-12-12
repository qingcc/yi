package main

import (
	"context"
	"flag"
	"github.com/qingcc/yi/utils/transutils"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:10003", "server address")
)
func main()  {
	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Trans", new(TransServer), "")
	s.Serve("tcp",*addr)
}


type TransServer struct {}

type Args struct {
	From string `json:"from"`
	To string `json:"to"`
	Query string `json:"query"`
	Ssl bool `json:"ssl"`
}

type Reply struct {
	Words string `json:"words"`
}

func (tr *TransServer)Trans(ctx context.Context, args Args, reply *Reply)  {
	res := transutils.Transfer(args.Query, args.From, args.To, args.Ssl)
	res = res
	return
}