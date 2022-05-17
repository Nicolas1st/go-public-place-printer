package jobs

import "net/http"

type JobsHandlers struct {
	SubmitJob func(w http.ResponseWriter, r *http.Request) error
	CancelJob func(w http.ResponseWriter, r *http.Request) error
}

func NewJobsHandlers(jobq jobqInterface, filer filerInterface) *JobsHandlers {
	jobsDependenices := &jobsDependencies{
		jobq:  jobq,
		filer: filer,
	}

	return &JobsHandlers{
		SubmitJob: jobsDependenices.SubmitJob,
		CancelJob: jobsDependenices.CancelJob,
	}
}
