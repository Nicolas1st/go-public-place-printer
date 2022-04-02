package fileuploader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFileHandler(c config) http.Handler {

	newFileName := newFileNamer(c)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// parsing form
		r.ParseForm()
		err := r.ParseMultipartForm(c.maxFileSize)

		// returning an error if parsing has failed
		if err != nil {
			response := newFileTooBigReponse(c.maxFileSize)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Get handler for filename, size and headers
		uploadedFile, handler, err := r.FormFile(c.formFieldFileName)
		if err != nil {
			response := newCouldNotReadFileResponse()
			json.NewEncoder(w).Encode(response)
			return
		}

		defer uploadedFile.Close()

		// info about the file
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Creating new file
		username := "user"
		fileOnServerName := newFileName(username)
		fileOnServer, err := os.Create(fileOnServerName)
		defer fileOnServer.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(fileOnServer, uploadedFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := newFileUploadedResponse(fileOnServerName)
		json.NewEncoder(w).Encode(response)
	})

}
