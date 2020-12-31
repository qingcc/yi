package main

import (
	"fmt"
	"github.com/qingcc/yi/lib/worker_pool"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

var wp = lib.NewWorkerPool(3)

type Demo struct {
	i int
}

func (d Demo) Do() {
	time.Sleep(time.Second)
	fmt.Println("done ", d.i)
}
func Decimal(value float64) float64 {
	// 只去浮点数的小数点后两位
	return math.Trunc(value*1e2+0.5) * 1e-2
}

func main() {
	num := (1 + 3.5/100) * 391
	logrus.Println("num:", Decimal(num))
	return
	for i := 0; i < 10000; i++ {
		d := Demo{i: i}
		wp.EnqueueJob(d)
		fmt.Println("enqueue i:", i)
	}
}
