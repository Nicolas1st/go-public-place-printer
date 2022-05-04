package worker

import (
	"fmt"
	"printer/persistence/model"
	"time"
)

type worker struct {
	jobq  jobqInterface
	filer filerInterface
}

func NewWorker(jobq jobqInterface, filer filerInterface) *worker {
	return &worker{
		jobq:  jobq,
		filer: filer,
	}
}

func (w *worker) ExecuteJob(job *model.Job) {
	time.Sleep(2 * time.Second)
	fmt.Println("Done executing the job")
}

// Start - sets up a printer to track the jobq state and perform
// printing whenever the job appears
func (w *worker) Start() {
	go func() {
		for {
			job := w.jobq.Dequeue()
			w.ExecuteJob(job)

			// move it out of worker later on
			w.filer.RemoveFile(job.FilePath)
		}
	}()
}
