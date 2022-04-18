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
		mainTemplateName:  "submit-file",
		templateFileNames: []string{"submit_file.html", "navbar.html"},
	})

	return &SubmitFilePage{
		tmpl: tmpl,
	}
}

func (page SubmitFilePage) Execute(w http.ResponseWriter, data any) error {
	return page.tmpl.Execute(w, data)
}
