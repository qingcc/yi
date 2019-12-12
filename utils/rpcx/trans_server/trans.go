package trans_server

import (
	"context"
	"github.com/qingcc/yi/utils/transutils"
)

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
