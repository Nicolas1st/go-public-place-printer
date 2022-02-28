package job

import "printer/interfaces"

type Job struct {
	ID   interfaces.JobID
	Func func() error
}

func NewJob(jobFunc func() error) Job {
	return Job{
		Func: jobFunc,
	}
}

func (job *Job) Execute() error {
	return job.Func()
}

func (job *Job) GetID() interfaces.JobID {
	return job.ID
}

func (job *Job) SetID(id interfaces.JobID) {
	job.ID = id
}
