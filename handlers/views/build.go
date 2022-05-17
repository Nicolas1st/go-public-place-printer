package views

import (
	"net/http"
	"printer/handlers/views/pages"
)

func buildView(p *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.Execute(w, pages.NewFlashMessages(), nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
