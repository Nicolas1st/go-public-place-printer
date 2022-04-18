package views

import (
	"net/http"
	"printer/web/pages"
)

func BuildLoginView(page *pages.LoginPage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := page.Execute(w, struct{ Greeting string }{Greeting: "Hello"})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
