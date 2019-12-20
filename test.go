package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/yi/utils"
	"github.com/qingcc/yi/utils/rpcx/trans_server/service"
	"github.com/tealeg/xlsx"
	"log"
	"strings"
	"time"
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

func main1()  {
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

func main()  {
//readfile("hotel.xlsx")
time.Sleep(time.Hour)
}

func readfile(file string) {
	// 打开文件
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		log.Println("打开文件失败！" + err.Error())
		return
	}
	if len(xlFile.Sheets) == 0 {
		log.Println("没有找到工作表")
	}

	unMap := make(map[string]bool)
	for k, row := range xlFile.Sheets[0].Rows {
		if k == 0 {
			continue
		}
		countryname := row.Cells[2].String()
		cen := strings.ToLower(row.Cells[3].String())
		if _, ok := unMap[countryname]; !ok {
			str := "\""+countryname+"\"" + ":\"" + cen+ "\","
			utils.Tracefile(str, "log.log")
			unMap[countryname] = true
		}
		//if _, ok := Country2Code[strings.ToLower(countryname)]; !ok {
		//	if _, has := unMap[countryname]; !has {
		//		unMap[countryname] = true
		//	}
		//}
	}
	fmt.Println("\n\nimport success")
	//for country, _ := range unMap {
	//	fmt.Println("\""+country +"\""+":"+"\"\"")
	//}
	fmt.Println(len(unMap))


}


