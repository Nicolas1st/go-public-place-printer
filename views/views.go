package views

import (
	"net/http"
	"printer/config"
	"printer/views/pages"
)

type Views struct {
	Login       http.HandlerFunc
	Signup      http.HandlerFunc
	SubmitFile  http.HandlerFunc
	UserManager http.HandlerFunc
}

func NewViews(htmlTemplatesPath string, endpoints config.Endpoints) *Views {
	pages := pages.NewPages(htmlTemplatesPath, endpoints)
	return &Views{
		Login:       BuildLoginView(pages.Login),
		Signup:      BuildSignupView(pages.Signup),
		SubmitFile:  BuildSubmitFileView(pages.SubmitFile),
		UserManager: BuildUserManagerView(pages.UserManager),
	}
}

func BuildLoginView(page *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := page.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func BuildSignupView(page *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := page.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func BuildSubmitFileView(page *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := page.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func BuildUserManagerView(page *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := page.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
