package pages

import (
	"fmt"
	"html/template"
	"path"
	"printer/handlers"
)

func buildPage(pageName string, pathToTemplates, templateToExecute string, templateNames ...string) *Page {
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

	return &Page{
		template:  template,
		Endpoints: handlers.DefaultEndpoints,
	}
}
