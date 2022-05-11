package pages

import (
	"html/template"
	"net/http"
	"printer/persistence/model"
)

type ProfilePage struct {
	tmpl *template.Template
}

type ProfilePageData struct {
	User   model.User
	Prints []model.Print
}

func NewProfilePage(htmlTemplatesPath string) *ProfilePage {
	tmpl := buildTemplate(htmlTemplatesPath, pageInfo{
		mainTemplateName:  "layout",
		templateFileNames: []string{"layout.html", "navbar.html", "private-links.html", "profile.html", "footer.html"},
	})

	return &ProfilePage{
		tmpl: tmpl,
	}
}

func (page ProfilePage) Execute(w http.ResponseWriter, data ProfilePageData) error {
	return page.tmpl.Execute(w, data)
}
