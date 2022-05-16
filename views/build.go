package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"printer/config"
)

type page struct {
	template  *template.Template
	Endpoints config.Endpoints
}

type pageData struct {
	Endpoints config.Endpoints
	Data      any
}

func buildPage(pageName string, endpoints config.Endpoints, pathToTemplates, templateToExecute string, templateNames ...string) *page {
	if len(templateNames) == 0 {
		panic("Can not build page with zerof files provided")
	}

	// prepend filepath
	withPaths := []string{}
	for _, fileName := range templateNames {
		withPaths = append(withPaths, path.Join(pathToTemplates, fileName))
	}

	// create template
	template := template.New(templateToExecute)
	template, err := template.ParseFiles(withPaths...)
	if err != nil {
		panic(fmt.Sprintf("Could not parse files for page %s", pageName))
	}

	return &page{
		template:  template,
		Endpoints: endpoints,
	}
}

func (p *page) execute(w http.ResponseWriter, data any) error {
	err := p.template.Execute(w, pageData{
		Endpoints: p.Endpoints,
		Data:      data,
	})

	return err
}
