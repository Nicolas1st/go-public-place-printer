package jobs

import (
	"io"
	"net/http"
	"printer/handlers"
	"printer/persistence/model"

	"github.com/gorilla/mux"
)

type filerInterface interface {
	StoreFile(uploadedFile io.Reader, username, submittedFilename string) (filepath string, err error)
	RemoveFile(filePath string) error
}

type jobqInterface interface {
	Enqueue(job *model.Job) model.JobID
	CancelJob(jobID model.JobID)
	GetAllJobs() []*model.Job
}

type jobsController struct {
	jobq      jobqInterface
	filer     filerInterface
	sessioner handlers.Sessioner
}

func NewApi(jobq jobqInterface, filer filerInterface, sessioner handlers.Sessioner) *mux.Router {
	c := &jobsController{
		jobq:      jobq,
		filer:     filer,
		sessioner: sessioner,
	}

	r := mux.NewRouter()

	r.HandleFunc(handlers.DefaultEndpoints.JobsApi, handlers.ForAdmin(c.sessioner, c.GetAllJobs)).Methods(http.MethodGet)
	r.HandleFunc(handlers.DefaultEndpoints.JobsApi, c.SubmitJob).Methods(http.MethodPost)
	r.HandleFunc(handlers.DefaultEndpoints.JobsApi+"/{id:[0-9]+}", c.CancelJob).Methods(http.MethodDelete)

	return r
}
