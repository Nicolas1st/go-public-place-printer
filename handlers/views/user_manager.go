package views

import (
	"net/http"
	"printer/handlers/views/pages"
)

func (c *viewsController) buildUserManagerView(p *pages.Pages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.Admin.UserManager.Execute(w, pages.NewFlashMessages(), nil)
	}
}
