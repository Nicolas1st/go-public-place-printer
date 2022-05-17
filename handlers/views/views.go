package views

import (
	"net/http"
	"printer/handlers"
)

type views struct {
	Login       http.HandlerFunc
	Signup      http.HandlerFunc
	SubmitFile  http.HandlerFunc
	UserManager http.HandlerFunc
}

func NewViews(htmlTemplatesPath string, endpoints handlers.Endpoints) *views {
	pages := newPages(htmlTemplatesPath, endpoints)
	return &views{
		Login:       buildView(pages.Login),
		Signup:      buildView(pages.Signup),
		SubmitFile:  buildView(pages.SubmitFile),
		UserManager: buildView(pages.UserManager),
	}
}

func buildView(p *page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
