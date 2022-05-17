package views

import (
	"net/http"
)

type views struct {
	Login       http.HandlerFunc
	SignUp      http.HandlerFunc
	SubmitFile  http.HandlerFunc
	UserManager http.HandlerFunc
}

func NewViews(htmlTemplatesPath string) *views {
	pages := newPages(htmlTemplatesPath)
	return &views{
		Login:       buildView(pages.Login),
		SignUp:      buildView(pages.Signup),
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
