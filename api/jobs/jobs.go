package jobs

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(jobq jobqInterface, filer filerInterface) *mux.Router {
	resourse := NewJobsResource(jobq, filer)

	r := mux.NewRouter()

	r.HandleFunc("/jobs",
		resourse.SubmitJob,
	).Methods(http.MethodPost)

	r.HandleFunc("/jobs",
		resourse.CancelJob,
	).Methods(http.MethodDelete)

	return r
}
