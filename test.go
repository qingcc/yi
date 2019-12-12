package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
)

var (
	addr = flag.String("addr", "localhost:10003", "server address")
)
func main()  {
	flag.Parse()
	gin.Default()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong, addr: " + *addr,
		})
	})
	r.Run(*addr) // 在 0.0.0.0:8080 上监听并服务
}



type Hello struct {
	Str string
}

func(h Hello) Run() {
	log.Println(h.Str)
}

func init() {
	log.Println("Starting...")

	c := cron.New()
	//h := Hello{"I Love You!"}
	// 添加定时任务
	//c.AddJob("*/2 * * * * * ", h)
	// 添加定时任务
	c.AddFunc("*/5 * * * * * ", func() {
		log.Println("hello word")
	})

	//s, err := cron.Parse("*/3 * * * * *")
	//if err != nil {
	//	log.Println("Parse error")
	//}
	//h2 := Hello{"I Hate You!"}
	//c.Schedule(s, h2)
	// 其中任务
	c.Start()
	// 关闭任务
	defer c.Stop()
	select {
	}
}