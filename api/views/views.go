package views

import (
	"net/http"
	"printer/web/pages"
)

type views struct {
	login      http.HandlerFunc
	signup     http.HandlerFunc
	submitFile http.HandlerFunc
}

func newViews(pages pages.Pages) *views {
	return &views{
		login:      BuildLoginView(pages.Login),
		signup:     BuildSignupView(pages.Signup),
		submitFile: BuildSubmitFileView(pages.SubmitFile),
	}
}

func NewRouter(pages pages.Pages) *http.ServeMux {
	views := newViews(pages)
	router := http.NewServeMux()

	router.HandleFunc("/login", views.login)
	router.HandleFunc("/signup", views.signup)
	router.HandleFunc("/submit-file", views.submitFile)

	return router
}
