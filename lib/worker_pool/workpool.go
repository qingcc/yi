package lib

import (
	"errors"
	"time"
)

type WorkPool struct {
	workCount   int
	workerQueue chan chan Job
	jobQueue    chan Job
}

func NewWorkerPool(workerCount int) (wp *WorkPool) {
	wp = &WorkPool{
		workCount:   workerCount,
		workerQueue: make(chan chan Job, workerCount),
		jobQueue:    make(chan Job),
	}
	wp.run()
	return
}

func (wp *WorkPool) run() {
	for i := 0; i < wp.workCount; i++ {
		w := newWorker(i)
		w.run(wp.workerQueue)
	}

	go func() {
		for {
			select {
			case job := <-wp.jobQueue:
				w := <-wp.workerQueue
				w <- job
			}
		}
	}()
}

func (wp *WorkPool) EnqueueJob(job Job) {
	wp.jobQueue <- job
}

//job必须在timeout时间内执行完，否则返回 job was abandoned， 并放弃执行该job
func (wp *WorkPool) EnqueueJobWithTimeOut(job Job, timeout time.Duration) (err error) {
	t := time.After(timeout)
	select {
	case wp.jobQueue <- job:
	case <-t:
		err = errors.New("job was abandoned")
	}
	return
}
