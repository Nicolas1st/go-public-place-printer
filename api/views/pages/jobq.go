package pages

import (
	"html/template"
	"net/http"
	"printer/persistence/model"
)

type JobqPage struct {
	tmpl *template.Template
}

func NewJobqPage(htmlTemplatesPath string) *JobqPage {
	tmpl := buildTemplate(htmlTemplatesPath, pageInfo{
		mainTemplateName:  "jobq",
		templateFileNames: []string{"jobq.html", "navbar.html"},
	})

	return &JobqPage{
		tmpl: tmpl,
	}
}

func (page JobqPage) Execute(w http.ResponseWriter, data []*model.Job) error {
	return page.tmpl.Execute(w, data)
}
