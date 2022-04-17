package jobs

import "printer/persistence/model"

type jobqInterface interface {
	Enqueue(job model.Job) model.JobID
	CancelJob(jobID model.JobID) error
}
