package jobs

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"printer/middlewares"
	"printer/persistence/model"
)

type jobsDependencies struct {
	jobq  jobqInterface
	filer filerInterface
}

func (resource *jobsDependencies) SubmitJob(w http.ResponseWriter, r *http.Request) error {
	// for now expecting the user the provide his name in the form
	var session *model.Session
	switch v := r.Context().Value(middlewares.ContextSessionKey).(type) {
	case *model.Session:
		session = v
	default:
		return errors.New("could not retrieve value from the context")
	}

	username := session.Username

	// extract the file form the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return errors.New("no file found in the submitted form")
	}

	// check whether the file is pdf
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("could not read the file")
	}
	mimeType := http.DetectContentType(bytes)
	if mimeType != "application/pdf" {
		return errors.New("wrong file format, the printer can only work with pdf files")
	}

	// storing the file
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return errors.New("could not read the file")
	}

	filepath, err := resource.filer.StoreFile(file, username, fileHeader.Filename)
	if err != nil {
		return errors.New("could not store the file")
	}

	// build job
	job := model.NewJob(filepath, fileHeader.Filename, username)
	resource.jobq.Enqueue(job)

	return nil
}

func (resource *jobsDependencies) CancelJob(w http.ResponseWriter, r *http.Request) error {
	var cancelJobRequest struct {
		ID model.JobID `json:"ID"`
	}
	err := json.NewDecoder(r.Body).Decode(&cancelJobRequest)
	if err != nil {
		return errors.New("bad request, try again")
	}

	resource.jobq.CancelJob(model.JobID(cancelJobRequest.ID))

	return nil
}
