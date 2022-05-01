package pages

import (
	"html/template"
	"net/http"
)

type SignupPage struct {
	tmpl *template.Template
}

func NewSignupPage(htmlTemplatesPath string) *SignupPage {
	tmpl := buildTemplate(htmlTemplatesPath, pageInfo{
		mainTemplateName:  "layout",
		templateFileNames: []string{"layout.html", "navbar.html", "public-links.html", "signup.html", "footer.html"},
	})

	return &SignupPage{
		tmpl: tmpl,
	}
}

func (page SignupPage) Execute(w http.ResponseWriter, data any) error {
	return page.tmpl.Execute(w, data)
}
