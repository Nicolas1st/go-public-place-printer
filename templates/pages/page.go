package pages

import (
	"fmt"
	"html/template"
	"path"
)

type Page struct {
	mainFileName string
	allFileNames []string
}

func (p *Page) BuildTemplate() *template.Template {
	template := template.New(p.GetMainFilePath())

	template, err := template.ParseFiles(p.GetAllFilePaths()...)
	if err != nil {
		fmt.Println(err)
	}

	return template
}

func GetAbsolutePath(filepath string) string {
	return path.Join("templates", "static", filepath)
}

func (p *Page) GetMainFilePath() string {
	return GetAbsolutePath(p.mainFileName)
}

func (p *Page) GetAllFilePaths() []string {
	var filepaths []string
	for _, name := range p.allFileNames {
		filepaths = append(filepaths, GetAbsolutePath(name))
	}

	return filepaths
}
