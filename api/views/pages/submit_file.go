package pages

import (
	"html/template"
	"net/http"
)

type SubmitFilePage struct {
	tmpl *template.Template
}

func NewSubmitFilePage(htmlTemplatesPath string) *SubmitFilePage {
	tmpl := buildTemplate(htmlTemplatesPath, pageInfo{
		mainTemplateName:  "layout",
		templateFileNames: []string{"layout.html", "navbar.html", "private-links.html", "submit-file.html", "footer.html"},
	})

	return &SubmitFilePage{
		tmpl: tmpl,
	}
}

func (page SubmitFilePage) Execute(w http.ResponseWriter, data any) error {
	return page.tmpl.Execute(w, data)
}
