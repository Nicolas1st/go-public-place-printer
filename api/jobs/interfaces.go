package jobs

import (
	"io"
	"printer/persistence/model"
)

type filerInterface interface {
	StoreFile(uploadedFile io.Reader, username, submittedFilename string) (filepath string, err error)
	RemoveFile(filePath string) error
}

type jobqInterface interface {
	Enqueue(job model.Job) model.JobID
	CancelJob(jobID model.JobID) error
}
