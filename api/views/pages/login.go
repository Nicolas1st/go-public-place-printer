package pages

import (
	"html/template"
	"net/http"
)

type LoginPage struct {
	tmpl *template.Template
}

func NewLoginPage(htmlTemplatesPath string) *LoginPage {
	tmpl := buildTemplate(htmlTemplatesPath, pageInfo{
		mainTemplateName:  "login.html",
		templateFileNames: []string{"login.html", "navbar.html"},
	})

	return &LoginPage{
		tmpl: tmpl,
	}
}

func (page LoginPage) Execute(w http.ResponseWriter, data any) error {
	return page.tmpl.Execute(w, data)
}
