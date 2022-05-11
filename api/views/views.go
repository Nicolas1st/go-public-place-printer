package views

import (
	"net/http"
	"printer/api/views/pages"
)

type Views struct {
	Login       http.HandlerFunc
	Signup      http.HandlerFunc
	SubmitFile  http.HandlerFunc
	UserManager http.HandlerFunc
	Profile     http.HandlerFunc
}

func NewViews(htmlTemplatesPath string, database DatabaseInterface) *Views {
	pages := pages.NewPages(htmlTemplatesPath)
	return &Views{
		Login:       BuildLoginView(pages.Login),
		Signup:      BuildSignupView(pages.Signup),
		SubmitFile:  BuildSubmitFileView(pages.SubmitFile),
		UserManager: BuildUserManagerView(pages.UserManager),
		Profile:     BuildProfileView(pages.Profile, database),
	}
}
