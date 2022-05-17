package pages

import (
	"html/template"
	"net/http"
	"printer/handlers"
)

type Page struct {
	template  *template.Template
	Endpoints handlers.Endpoints
}

type pageData struct {
	Endpoints handlers.Endpoints
	Data      any
}

func (p *Page) Execute(w http.ResponseWriter, data any) error {
	err := p.template.Execute(w, pageData{
		Endpoints: handlers.DefaultEndpoints,
		Data:      data,
	})

	return err
}
