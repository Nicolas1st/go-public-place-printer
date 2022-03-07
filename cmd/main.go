package main

import (
	"fmt"
	"net/http"
	"printer/api/jobs"
	"printer/interfaces"
	"printer/job"
	"printer/jobqueue"
)

func main() {
	q := jobqueue.NewJobQueue()

	createJob := func() interfaces.Job {
		return &job.Job{}
	}

	jobsApi := jobs.NewRouter(q, createJob)
	fmt.Println("The server starting on localhost:7777")
	http.ListenAndServe(":7777", jobsApi)
}
