package job_worder_pool

type WorkPool struct {
	workerCount int
	JobQueue chan Job
	WorderQueue chan chan Job
}

func NewWorkPool(workerCount int) *WorkPool {
	wp := &WorkPool{
		workerCount:   workerCount,
		JobQueue:    make(chan Job),
		WorderQueue: make(chan chan Job, workerCount),
	}
	wp.run()

	return wp
}

func (wp *WorkPool)run()  {
	for i:=0 ; i < wp.workerCount ; i++ {
		work := newWorker(i)
		work.Run(wp.WorderQueue)
	}

	go func() {
		for {
			select {
			case job := <- wp.JobQueue:
				worker := <-wp.WorderQueue
				worker <- job
			}
		}
	}()
}

func (wp *WorkPool)EnqueueJob(job Job)  {
	wp.JobQueue <- job
}