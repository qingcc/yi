package service

import (
	trans_service "github.com/qingcc/yi/commobj/utils"
	commobj_trans "github.com/qingcc/yi/commobj/utils/rpcx"
	"github.com/qingcc/yi/utils"
	"log"
	"testing"
)

func TestTransfer(t *testing.T) {
	req := trans_service.Args{
		From:  "zh",
		To:    "en",
		Query: "文档",
		Ssl:   false,
	}
	res := &commobj_trans.TransResponse{}
	Transfer(req, res)
	log.Println("res:", utils.ToJson(res), "dst:", res.TransResult.Dst)
}
