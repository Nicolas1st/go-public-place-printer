package views

import (
	"net/http"
	"printer/web/pages"
)

func BuildSubmitFileView(page *pages.SubmitFilePage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := page.Execute(w, struct{ Greeting string }{Greeting: "Hello"})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
