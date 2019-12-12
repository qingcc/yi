package job_worder_pool

type Job interface {
	Do()
}

type worker struct {
	Index int
	JobQueue chan Job
}

func newWorker(index int) worker {
	return worker{Index:index, JobQueue:make(chan Job)}
}

func (w worker)Run(wq chan chan Job)  {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <- w.JobQueue:
				job.Do()
			}
		}
	}()
}
