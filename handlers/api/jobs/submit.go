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

	// проверить на возможность печатиo
	u, _ := c.logger.GetUserByID(session.UserID)
	if !u.CanUsePrinter {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Возможность печати заблокирована")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// extract the file form the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Необходимо выбрать файл для печати")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// check whether the file is pdf
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Не получилось открыть файл")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// проверить тип документа
	mimeType := http.DetectContentType(bytes)
	if mimeType != "application/pdf" {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Возможна печать только PDF файлов")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// сохранить файл
	file.Seek(0, io.SeekStart)
	filename := strings.ReplaceAll(fileHeader.Filename, " ", "")
	filepath, err := c.filer.StoreFile(file, session.Username, filename)
	if err != nil {
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, "Невозможно сохранить файл, закончилась память")
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// узнать количество страниц в документе
	cmd := fmt.Sprintf("pdfinfo %v | awk '/^Pages:/ {print $2}'", filepath)
	command := exec.Command("bash", "-c", cmd)
	out, _ := command.CombinedOutput()
	numberOfPages, _ := strconv.Atoi(strings.Replace(string(out), "\n", "", -1))

	// проверить не превзойден ли лимит
	pageLimit, _ := c.logger.GetPageLimit(session.UserID)
	printedOverTheLastMonth := c.logger.GetNumberOfPagesPrintedByUserDuringTheLastMonth(session.UserID)
	if pageLimit < uint(numberOfPages)+uint(printedOverTheLastMonth) {
		pagesLeft := int(pageLimit) - printedOverTheLastMonth
		message := fmt.Sprintf("Доступно страниц %v, в документе %v", pagesLeft, numberOfPages)
		jsonResponse.FlashMessages = append(jsonResponse.FlashMessages, message)
		json.NewEncoder(w).Encode(&jsonResponse)
		return
	}

	// сохранить лог о количестве страниц в документе
	c.logger.SavePrint(*session.User, filename, filepath, numberOfPages)

	// создать задачу
	job := model.NewJob(filepath, fileHeader.Filename, session.User)
	jobID := c.jobq.Push(job)

	jsonResponse.Success = true
	jsonResponse.JobID = string(jobID)
	jsonResponse.Filename = fileHeader.Filename

	json.NewEncoder(w).Encode(&jsonResponse)
}
