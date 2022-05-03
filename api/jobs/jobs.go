package jobs

import "net/http"

func NewRouter(jobq jobqInterface, filer filerInterface) *http.ServeMux {
	router := http.NewServeMux()

	resource := NewJobsResource(jobq, filer)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			resource.SubmitJob(w, r)
		case http.MethodDelete:
			resource.CancelJob(w, r)
		}
	})

	return router
}
