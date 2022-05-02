package views

import (
	"net/http"
	"printer/api/views/pages"
)

func NewAdminViews(pages *pages.Pages, jobq jobqInterface) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/jobq", BuildJobqView(pages.Jobq, jobq))

	return router
}

func NewPublicViews(pages *pages.Pages) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/login", BuildLoginView(pages.Login))
	router.HandleFunc("/signup", BuildSignupView(pages.Signup))

	return router
}

func NewPrivateViews(pages *pages.Pages) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/submit-file", BuildSubmitFileView(pages.SubmitFile))

	return router
}
