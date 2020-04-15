package lib

type Job interface {
	Do()
}

type worker struct {
	index    int
	jobQueue chan Job
}

func newWorker(i int) *worker {
	return &worker{index: i, jobQueue: make(chan Job)}
}

func (w *worker) run(wq chan chan Job) {
	go func() {
		for {
			wq <- w.jobQueue
			select {
			case job := <-w.jobQueue:
				job.Do()
			}
		}

	}()
}
