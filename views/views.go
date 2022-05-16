package views

import (
	"net/http"
	"printer/config"
)

type views struct {
	Login       http.HandlerFunc
	Signup      http.HandlerFunc
	SubmitFile  http.HandlerFunc
	UserManager http.HandlerFunc
}

func NewViews(htmlTemplatesPath string, endpoints config.Endpoints) *views {
	pages := newPages(htmlTemplatesPath, endpoints)
	return &views{
		Login:       buildLoginView(pages.Login),
		Signup:      buildSignupView(pages.Signup),
		SubmitFile:  buildSubmitFileView(pages.SubmitFile),
		UserManager: buildUserManagerView(pages.UserManager),
	}
}

func buildLoginView(p *page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func buildSignupView(p *page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func buildSubmitFileView(p *page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func buildUserManagerView(p *page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
