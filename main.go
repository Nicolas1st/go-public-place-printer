package main

import (
	"net/http"
	"printer/api/jobs"
	"printer/api/views"
	"printer/api/views/pages"
	"printer/persistence/filer"
	"printer/persistence/jobq"
)

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:8880",
	}

	jobq := jobq.NewJobQueue()
	filer := filer.NewFiler("./files", 2<<30)

	views := views.NewRouter(*pages.NewPages("./web/html"), jobq)
	http.Handle("/", views)

	jobsApi := jobs.NewRouter(jobq, filer)
	http.Handle("/jobs", jobsApi)

	server.ListenAndServe()
}
