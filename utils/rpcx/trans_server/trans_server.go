package main

import (
	"context"
	"flag"
	commobj "github.com/qingcc/yi/commobj/utils"
	"github.com/qingcc/yi/utils/rpcx/trans_server/service"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:10003", "server address")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	if err := s.RegisterName("TransServer", new(TransServer), ""); err != nil {
		panic(err)
	}
	s.Serve("tcp", *addr)
}

type TransServer int

func (tr *TransServer) Transfer(ctx context.Context, args commobj.Args, reply *commobj.Reply) error {
	//reply.Query = service.Transfer(args.Query, args.From, args.To, args.Ssl)
	reply.Query = service.Trans(args.Query, args.From, args.To, args.Ssl)
	return nil
}
