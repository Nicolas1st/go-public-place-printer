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
		mainTemplateName:  "layout",
		templateFileNames: []string{"layout.html", "navbar.html", "public-links.html", "login.html", "footer.html"},
	})

	return &LoginPage{
		tmpl: tmpl,
	}
}

func (page LoginPage) Execute(w http.ResponseWriter, data any) error {
	return page.tmpl.Execute(w, data)
}
