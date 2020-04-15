package cli

import (
	commobj "github.com/qingcc/yi/commobj/utils"
	"github.com/qingcc/yi/utils"
	"log"
	"testing"
)

func TestGetTransServiceClient(t *testing.T) {
	args := commobj.Args{
		From:  "zh",
		To:    "en",
		Query: "中文",
		Ssl:   true,
	}
	reply := GetTransServiceClient().Trans(args)
	log.Println("reply:", utils.ToJson(reply))
}
