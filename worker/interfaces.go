package worker

import "printer/persistence/model"

type jobqInterface interface {
	Dequeue() *model.Job
}

type filerInterface interface {
	RemoveFile(filePath string) error
}
