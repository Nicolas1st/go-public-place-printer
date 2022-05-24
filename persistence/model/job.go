package model

type JobStatus int

type JobID string

type Job struct {
	ID                JobID  `json:"id"`
	SubmittedFileName string `json:"filename"`
	StoredFilePath    string
	User              *User
}

func NewJob(storedFilePath, submittedFileName string, user *User) *Job {
	return &Job{
		StoredFilePath:    storedFilePath,
		SubmittedFileName: submittedFileName,
		User:              user,
	}
}
