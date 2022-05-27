package views

import (
	"net/http"
	"printer/handlers/views/pages"
)

func (c *viewsController) buildFileRemovedView(p *pages.Pages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.Admin.FileRemoved.Execute(w, pages.NewFlashMessages(), nil)
	}
}
