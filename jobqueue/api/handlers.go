package jobs

import (
	"encoding/json"
	"net/http"
	"printer/interfaces"
)

func BuildSubmitJob(q interfaces.JobQueue, createJob func() interfaces.Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// build job
		job := createJob()
		jobID := q.Enqueue(job)

		// create response
		responseBody := NewResponse(jobID)
		json.NewEncoder(w).Encode(responseBody)
	}
}

func BuildCancelJob(q interfaces.JobQueue) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// read the request body
		body := CancelJobRequest{}
		json.NewDecoder(r.Body).Decode(&body)

		err := q.CancelJob(body.ID)
		if err == nil {
			json.NewEncoder(w).Encode("Job has been canceled")
		} else {
			json.NewEncoder(w).Encode("Job with the id provided does not exist")
		}

	}
}
