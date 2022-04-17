package model

type JobID uint

type Job struct {
	ID       JobID
	FilePath string
	Username string
}

func NewJob(id JobID, filepath, username string) *Job {
	return &Job{
		ID:       id,
		FilePath: filepath,
		Username: username,
	}
}
