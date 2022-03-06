package templates

import (
	"fmt"
	"html/template"
	"printer/templates/pages"
)

type Pages struct {
}

type Templates struct {
	Signup *template.Template
	Signin *template.Template
}

func BuildTemplate(p *pages.Page) *template.Template {
	template := template.New(p.GetMainFilePath())

	template, err := template.ParseFiles(p.GetAllFilePaths()...)
	if err != nil {
		fmt.Println(err)
	}

	return template
}

func NewTemplates() *Templates {
	return &Templates{
		Signup: pages.SignUpPage.BuildTemplate(),
		Signin: pages.SignInPage.BuildTemplate(),
	}
}
