package pages

import (
	"html/template"
	"log"
	"path"
)

type Page struct {
	mainFileName string
	allFileNames []string
}

func (p *Page) BuildTemplate() *template.Template {
	mainFilePath := getPathToTemplate(p.mainFileName)
	allFilePaths := []string{}
	for _, name := range p.allFileNames {
		allFilePaths = append(allFilePaths, getPathToTemplate(name))
	}

	template, err := template.New(mainFilePath).ParseFiles(allFilePaths...)
	if err != nil {
		log.Fatal(err)
	}

	return template
}

func getPathToTemplate(templateFileName string) string {
	return path.Join("templates", "static", templateFileName)
}
