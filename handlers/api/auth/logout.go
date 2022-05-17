package auth

import (
	"net/http"
	"printer/handlers"
)

func (c *authController) Logout(w http.ResponseWriter, r *http.Request) {
	if authCookie, ok := handlers.GetAuthCookie(r); ok {
		c.sessions.RemoveSession(authCookie.Value)
		handlers.RemoveAuthCookie(w)
	}

	http.Redirect(w, r, handlers.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}
