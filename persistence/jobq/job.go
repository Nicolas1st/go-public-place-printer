package jobq

type JobStatus int

const (
	StatusCancelled = iota
	StatusToBeDone
)

type JobID uint

type Job struct {
	ID                JobID
	SubmittedFileName string
	StoredFilePath    string
	Username          string
	Status            JobStatus
}

func NewJob(storedFilePath, submittedFileName, username string) *Job {
	return &Job{
		StoredFilePath:    storedFilePath,
		SubmittedFileName: submittedFileName,
		Username:          username,
		Status:            StatusToBeDone,
	}
}

func (j *Job) CancelJob() {
	j.Status = StatusCancelled
}

func (j *Job) SetID(id JobID) {
	j.ID = id
}
