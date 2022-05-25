package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"printer/handlers"
	"printer/persistence/model"
	"strconv"
	"strings"
)

type SubmitJobResponseSchema struct {
	Success       bool     `json:"success"`
	JobID         string   `json:"jobID"`
	Filename      string   `json:"filename"`
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
	filename := strings.ReplaceAll(fileHeader.Filename, " ", "")
	filepath, err := c.filer.StoreFile(file, session.Username, filename)
	if err != nil {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Could not store the file, ran out of memory")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// сохранить лог о количестве страниц в документе
	cmd := fmt.Sprintf("pdfinfo %v | awk '/^Pages:/ {print $2}'", filename)
	command := exec.Command("bash", "-c", cmd)
	out, _ := command.CombinedOutput()
	numberOfPages, _ := strconv.Atoi(string(out))
	c.logger.SavePrint(*session.User, filename, numberOfPages)

	// build job
	job := model.NewJob(filepath, fileHeader.Filename, session.User)
	jobID := c.jobq.Push(job)

	jsonResponse.Success = true
	jsonResponse.JobID = string(jobID)
	jsonResponse.Filename = fileHeader.Filename

	json.NewEncoder(w).Encode(&jsonResponse)
}
