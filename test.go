package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/yi/utils/transutils"
)

var (
	addr = flag.String("addr", "localhost:10003", "server address")
)
func main()  {
	flag.Parse()
	res := transutils.Transfer("中文", "zh", "en", false)
}

func f()  {
	gin.Default()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong, addr: " + *addr,
		})
	})
	r.Run(*addr) // 在 0.0.0.0:8080 上监听并服务
}