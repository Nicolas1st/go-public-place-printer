package jobs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"printer/persistence/model"
)

type jobsResource struct {
	jobq  jobqInterface
	filer filerInterface
}

func NewJobsResource(jobq jobqInterface, filer filerInterface) *jobsResource {
	return &jobsResource{
		jobq:  jobq,
		filer: filer,
	}
}

func (resource *jobsResource) SubmitJob(w http.ResponseWriter, r *http.Request) {
	// for now expecting the user the provide his name in the form
	username := r.PostFormValue("username")
	// extract the file form the form
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			ErrorText string
		}{
			ErrorText: "Could not extract the file from the form",
		})
		return
	}

	// storing the file
	filepath, err := resource.filer.StoreFile(file, username, fileHeader.Filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			ErrorText string
		}{
			ErrorText: "Could not save the file, the error has been logged, try renaming the file",
		})
		return
	}

	// build job
	job := model.NewJob(filepath, username)
	jobID := resource.jobq.Enqueue(*job)

	// create response
	responseBody := newResponse(jobID)
	json.NewEncoder(w).Encode(responseBody)
}

func (resouce *jobsResource) CancelJob(w http.ResponseWriter, r *http.Request) {
	// parsing the request
	request := cancelJobRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			ErrorText string
		}{
			ErrorText: "Could not parse the json the body of the request",
		})
		return
	}

	err = resouce.jobq.CancelJob(model.JobID(request.ID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			ErrorText string
		}{
			ErrorText: "Could not parse the json the body of the request",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct {
			ErrorText string
		}{
			ErrorText: fmt.Sprint(err),
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Message string
		}{
			Message: "The job has been cancelled",
		})
	}
}
