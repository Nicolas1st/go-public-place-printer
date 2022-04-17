package model

type JobID uint

type Job struct {
	ID       JobID
	FilePath string
	Username string
}

func NewJob(filepath, username string) *Job {
	return &Job{
		FilePath: filepath,
		Username: username,
	}
}

func (j *Job) SetID(id JobID) {
	j.ID = id
}
