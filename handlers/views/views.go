package views

import (
	"net/http"
	"printer/handlers/views/pages"
)

type views struct {
	Login       http.HandlerFunc
	SignUp      http.HandlerFunc
	SubmitFile  http.HandlerFunc
	UserManager http.HandlerFunc
}

func NewViews(htmlTemplatesPath string) *views {
	pages := pages.NewPages(htmlTemplatesPath)
	return &views{
		Login:       buildView(pages.Login),
		SignUp:      buildView(pages.Signup),
		SubmitFile:  buildView(pages.SubmitFile),
		UserManager: buildView(pages.UserManager),
	}
}

func buildView(p *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
