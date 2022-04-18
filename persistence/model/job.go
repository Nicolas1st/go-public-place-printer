package model

type JobStatus int

const (
	StatusCancelled = iota
	StatusToBeDone
)

type JobID uint

type Job struct {
	ID       JobID
	FilePath string
	Username string
	Status   JobStatus
}

func NewJob(filepath, username string) *Job {
	return &Job{
		FilePath: filepath,
		Username: username,
		Status:   StatusToBeDone,
	}
}

func (j *Job) CancelJob() {
	j.Status = StatusCancelled
}

func (j *Job) SetID(id JobID) {
	j.ID = id
}
