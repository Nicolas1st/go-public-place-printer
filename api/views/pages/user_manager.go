package pages

import (
	"html/template"
	"net/http"
)

type UserManagerPage struct {
	tmpl *template.Template
}

func NewUserManagerPage(htmlTemplatesPath string) *UserManagerPage {
	tmpl := buildTemplate(htmlTemplatesPath, pageInfo{
		mainTemplateName:  "layout",
		templateFileNames: []string{"layout.html", "navbar.html", "admin-links.html", "user-manager.html", "footer.html"},
	})

	return &UserManagerPage{
		tmpl: tmpl,
	}
}

func (page UserManagerPage) Execute(w http.ResponseWriter) error {
	return page.tmpl.Execute(w, nil)
}
