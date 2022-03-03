package jobs

import (
	"net/http"
	"printer/interfaces"

	"github.com/gorilla/mux"
)

func NewRouter(q interfaces.JobQueue, createJob func() interfaces.Job) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/",
		BuildSubmitJob(q, createJob),
	).Methods(http.MethodPost)

	r.HandleFunc("/",
		BuildCancelJob(q),
	).Methods(http.MethodDelete)

	return r
}
