package views

import (
	"net/http"
	"printer/api/views/pages"
)

type Views struct {
	Login      http.HandlerFunc
	Signup     http.HandlerFunc
	SubmitFile http.HandlerFunc
}

func NewViews(htmlTemplatesPath string) *Views {
	pages := pages.NewPages(htmlTemplatesPath)
	return &Views{
		Login:      BuildLoginView(pages.Login),
		Signup:     BuildSignupView(pages.Signup),
		SubmitFile: BuildSubmitFileView(pages.SubmitFile),
	}
}
