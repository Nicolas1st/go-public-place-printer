package views

import (
	"net/http"
	"printer/api/views/pages"
)

func BuildSubmitFileView(page *pages.SubmitFilePage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := page.Execute(w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
