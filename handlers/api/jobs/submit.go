package jobs

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"printer/handlers"
	"printer/persistence/model"
)

type SubmitJobResponseSchema struct {
	FlashMessages []string `json:"flashMessages"`
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
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "No file provided")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// check whether the file is pdf
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Corrupted file")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	mimeType := http.DetectContentType(bytes)
	if mimeType != "application/pdf" {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "The file provided is not a pdf file")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// storing the file
	file.Seek(0, io.SeekStart)

	filepath, err := c.filer.StoreFile(file, session.Username, fileHeader.Filename)
	if err != nil {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Could not store the file, ran out of memory")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// build job
	job := model.NewJob(filepath, fileHeader.Filename, session.Username)
	c.jobq.Enqueue(job)

	json.NewEncoder(w).Encode(&jsonResponse)
}
