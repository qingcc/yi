package main

import (
	"fmt"
	"github.com/qingcc/yi/lib/worker_pool"
)

var wp = lib.NewWorkerPool(3)

type Demo struct {
	i int
}

func (d Demo) Do() {
	fmt.Println(d.i)
}

func main() {
	for i := 0; i < 10; i++ {
		d := Demo{i: i}
		wp.EnqueueJob(d)
	}
}
