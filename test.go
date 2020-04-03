package main

import (
	"context"
	"fmt"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func getHotelIds(file string) {
	// 打开文件
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		log.Println("打开文件失败！" + err.Error())
		return
	}
	if len(xlFile.Sheets) == 0 {
		log.Println("没有找到工作表")
	}
	for k, row := range xlFile.Sheets[0].Rows {
		if k == 0 {
			continue
		}
		row = row
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func init() {
	t := time.Now().UTC()
	date := fmt.Sprintf("%s UTC", t.Format("Mon,02 Jan 2006 15:04:05"))
	fmt.Printf(date)
	zap.AddCaller()
}

//获取系统环境变量
func getEnv() {
	switch os.Getenv("GOPATH") {
	case "go":
	default:
		log.Println("[warn] ...")
		return
	}
}

func init() {
	orderInfoLists := []int{1, 2, 3, 4, 5, 55, 6, 7, 8, 9, 0}
	for _, v := range orderInfoLists {
		go accordingToBookDateUpdateOrder(v)
	}
}

func accordingToBookDateUpdateOrder(i int) {
	log.Println("i:", i)
}
