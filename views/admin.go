package views

import (
	"net/http"
	"printer/api/views/pages"
	"printer/persistence/model"
)

type jobqInterface interface {
	GetAllJobs() []*model.Job
}

func BuildJobqView(page *pages.JobqPage, jobq jobqInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobs := jobq.GetAllJobs()
		err := page.Execute(w, jobs)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func BuildUserManagerView(page *pages.UserManagerPage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := page.Execute(w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
