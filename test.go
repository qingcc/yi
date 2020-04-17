package main

import (
	"context"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
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
