package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/yi/utils/rpcx/trans_server/service"
)

var (
	addr = flag.String("addr", "localhost:10003", "server address")
)
func f()  {
	flag.Parse()
	//res := service.Transfer("中文", "zh", "en", false)
	res := service.Trans("中文", "zh", "en", false)
	fmt.Println(res)
}

func main()  {
	flag.Parse()
	gin.Default()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		res := service.Trans("中文", "zh", "en", false)
		fmt.Println(res)
		service.Transfer("中文", "zh", "en", false)
		c.JSON(200, gin.H{
			"message": "pong, addr: " + *addr,
		})
	})
	r.Run(*addr) // 在 0.0.0.0:8080 上监听并服务
}