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
	Endpoints     handlers.Endpoints
	FlashMessages []flashMessage
	Data          any
}

func (p *Page) Execute(w http.ResponseWriter, flash *FlashMessages, data any) error {
	err := p.template.Execute(w, pageData{
		Endpoints:     handlers.DefaultEndpoints,
		FlashMessages: flash.messages,
		Data:          data,
	})

	return err
}
