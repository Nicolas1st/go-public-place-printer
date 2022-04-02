package main

import (
	"fmt"
	"net/http"

	"printer/api/jobs"
	"printer/interfaces"
	"printer/job"
	"printer/jobqueue"
	"printer/pages"
	"printer/views"
)

var host = "localhost:7777"

func main() {
	q := jobqueue.NewJobQueue()
	jobsApi := jobs.NewRouter(q,
		func() interfaces.Job {
			return &job.Job{}
		},
	)

	templates := pages.NewTemplates("../static/html")
	views := views.NewRouter(templates)

	http.Handle("/jobs", jobsApi)
	http.Handle("/pages", views)

	fmt.Printf("The server has start on %v\n", host)
	http.ListenAndServe(host, views)
}
