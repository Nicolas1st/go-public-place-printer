package jobs

import (
	"io"
	"io/ioutil"
	"net/http"
	"printer/handlers"
	"printer/persistence/model"
)

type SubmitJobResponseSchema struct {
	RedirectionURL string   `json:"redirectionURL"`
	ErrorMessages  []string `json:"flashMessage"`
}

func (c *jobsController) SubmitJob(w http.ResponseWriter, r *http.Request) {
	session, ok := handlers.GetSession(c.sessioner, r)
	if !ok {
		http.Redirect(w, r, handlers.DefaultEndpoints.LoginPage, http.StatusSeeOther)
		return
	}

	var jsonResponse SubmitJobResponseSchema

	// extract the file form the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		jsonResponse.ErrorMessages = append(jsonResponse.ErrorMessages, "No file provided")
		return
	}

	// check whether the file is pdf
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		jsonResponse.ErrorMessages = append(jsonResponse.ErrorMessages, "Corrupted file")
		return
	}

	mimeType := http.DetectContentType(bytes)
	if mimeType != "application/pdf" {
		jsonResponse.ErrorMessages = append(jsonResponse.ErrorMessages, "The file provided is not a pdf file")
		return
	}

	// storing the file
	file.Seek(0, io.SeekStart)

	filepath, err := c.filer.StoreFile(file, session.Username, fileHeader.Filename)
	if err != nil {
		jsonResponse.ErrorMessages = append(jsonResponse.ErrorMessages, "Could not store the file, ran out of memory")
		return
	}

	// build job
	job := model.NewJob(filepath, fileHeader.Filename, session.Username)
	c.jobq.Enqueue(job)
}
