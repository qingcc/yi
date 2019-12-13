package main

import (
	"flag"
	"testing"
)

var (
	a     = flag.String("addr", "tcp@localhost:10003", "service address")
	query = flag.String("q", "需要翻译的语句", "query text")
	from  = flag.String("q", "zh", "from language")
	to    = flag.String("q", "en", "to language")
	ssl   = flag.Bool("s", false, "ssl")
)

func TestTransServer_Trans(t *testing.T) {
	flag.Parse()
	//s := client.NewPeer2PeerDiscovery("tcp@" + *a, "")

}
