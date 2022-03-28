package pages

import (
	"fmt"
	"html/template"
	"log"
	"path"
)

type Page struct {
	mainTemplateName string
	allTemplateNames []string
}

func (p *Page) BuildTemplate(htmlTemplatesPath string) *template.Template {
	tmpl := template.New(p.mainTemplateName)

	allFilePaths := []string{}
	for _, name := range p.allTemplateNames {
		allFilePaths = append(allFilePaths, getPathToTemplate(htmlTemplatesPath, name))
	}

	template, err := tmpl.ParseFiles(allFilePaths...)
	if err != nil {
		// The project should not start if the templates can not be built
		fmt.Println("Make you sure the path to the html templates was correctly specified")
		fmt.Println("The path should be relative to the main executable")
		log.Fatal(err)
	}

	return template
}

func getPathToTemplate(htmlTemplatesPath, templateName string) string {
	return path.Join(htmlTemplatesPath, templateName)
}
