package main

import (
	"context"
	"flag"
	commobj "github.com/qingcc/yi/commobj/utils"
	"github.com/smallnest/rpcx/client"
	"log"
	"testing"
)

var (
	a     = flag.String("addr1", "47.112.210.86:10011", "service address")
	query = flag.String("q", "需要翻译的语句", "query text")
	from  = flag.String("f", "zh", "from language")
	to    = flag.String("t", "en", "to language")
	ssl   = flag.Bool("s", false, "ssl")
)

func TestTransServer_Trans(t *testing.T) {
	flag.Parse()
	d := client.NewPeer2PeerDiscovery("tcp@"+*a, "")
	xclient := client.NewXClient("TransServer", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	for {
		args := commobj.Args{From: *from, To: *to, Query: *query, Ssl: *ssl}
		reply := &commobj.Reply{}
		xclient.Call(context.Background(), "Transfer", args, reply)
		log.Println(reply.Query)
	}
}
