package lib

import (
	"github.com/qingcc/yi/utils"
	"go.uber.org/ratelimit"
	"log"
	"sync"
)

type WorkPool struct {
	WorkerCount int
	Tasks []interface{}
	DoFunc func(interface{})
	RateLimitCount int //per second, per work
	IsOnce bool //true when tasks is null escape
	Done chan bool
}

func (wp *WorkPool)work(wg *sync.WaitGroup, taskCh chan interface{})  {
	utils.RecoverPanic(true)
	defer wg.Done()

	var r1 ratelimit.Limiter
	if wp.RateLimitCount > 0 {
		r1 = ratelimit.New(wp.RateLimitCount)
	}else {
		r1 = ratelimit.NewUnlimited()
	}
	for {
		task := <-taskCh
		r1.Take()
		wp.DoFunc(task)
		if wp.IsOnce && task == wp.Tasks[len(wp.Tasks)-1] {
			close(wp.Done)
		}
	}
}

func (wp *WorkPool)Process()  {
	var wg sync.WaitGroup
	taskCh := make(chan interface{})
	defer close(taskCh)

	go func() {
		log.Println("start executing task, task count:", len(wp.Tasks))
		for i, t := range wp.Tasks {
			log.Printf("enqueue job %d", i)
			taskCh<-t
		}
	}()

	wg.Add(wp.WorkerCount)

	log.Printf("worker count %d", wp.WorkerCount)
	for i:=0 ; i<wp.WorkerCount ; i++ {
		go wp.work(&wg, taskCh)
	}
	wg.Wait()
}