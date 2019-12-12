package trans_server

import (
	"flag"
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
