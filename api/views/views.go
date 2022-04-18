package views

import (
	"net/http"
	"printer/api/views/pages"
)

type views struct {
	jobq       http.HandlerFunc
	login      http.HandlerFunc
	signup     http.HandlerFunc
	submitFile http.HandlerFunc
}

func newViews(pages pages.Pages, jobq jobqInterface) *views {
	return &views{
		jobq:       BuildJobqView(pages.Jobq, jobq),
		login:      BuildLoginView(pages.Login),
		signup:     BuildSignupView(pages.Signup),
		submitFile: BuildSubmitFileView(pages.SubmitFile),
	}
}

func NewRouter(pages pages.Pages, jobq jobqInterface) *http.ServeMux {
	views := newViews(pages, jobq)
	router := http.NewServeMux()

	router.HandleFunc("/login", views.login)
	router.HandleFunc("/signup", views.signup)
	router.HandleFunc("/submit-file", views.submitFile)
	router.HandleFunc("/jobq", views.jobq)

	return router
}
