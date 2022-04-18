package views

import "printer/persistence/model"

type jobqInterface interface {
	GetAllJobs() []*model.Job
}
