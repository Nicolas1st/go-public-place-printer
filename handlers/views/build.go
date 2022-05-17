package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"printer/handlers"
)

type page struct {
	template  *template.Template
	Endpoints handlers.Endpoints
}

type pageData struct {
	Endpoints handlers.Endpoints
	Data      any
}

func buildPage(pageName string, pathToTemplates, templateToExecute string, templateNames ...string) *page {
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
		// communicate what is wrong with templates
		fmt.Println(err)
		panic(fmt.Sprintf("Could not parse files for page %s", pageName))
	}

	return &page{
		template:  template,
		Endpoints: handlers.DefaultEndpoints,
	}
}

func (p *page) execute(w http.ResponseWriter, data any) error {
	err := p.template.Execute(w, pageData{
		Endpoints: handlers.DefaultEndpoints,
		Data:      data,
	})

	return err
}
