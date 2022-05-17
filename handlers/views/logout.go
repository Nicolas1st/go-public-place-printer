package views

import (
	"net/http"
	"printer/handlers"
)

func (c *viewsController) Logout(w http.ResponseWriter, r *http.Request) {
	if authCookie, ok := handlers.GetAuthCookie(r); ok {
		c.sessioner.RemoveSession(authCookie.Value)
		handlers.RemoveAuthCookie(w)
	}

	http.Redirect(w, r, handlers.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}
