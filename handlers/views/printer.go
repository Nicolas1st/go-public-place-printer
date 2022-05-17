package views

import (
	"net/http"
	"printer/handlers"
	"printer/handlers/views/pages"
)

func (c *viewsController) buildPrinterView(p *pages.Pages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, ok := handlers.GetSession(c.sessioner, r)
		if !ok {
			http.Redirect(w, r, handlers.DefaultEndpoints.LoginPage, http.StatusSeeOther)
		}

		if session.User.IsAdmin() {
			p.Admin.Printer.Execute(w, pages.NewFlashMessages(), nil)
		} else {
			p.Private.Printer.Execute(w, pages.NewFlashMessages(), nil)
		}
	}
}
