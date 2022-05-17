package auth

import (
	"net/http"
	"printer/config"
	"printer/pkg/cookie"
)

func (c *authController) Logout(w http.ResponseWriter, r *http.Request) {
	if authCookie, ok := cookie.GetAuthCookie(r); ok {
		c.sessions.RemoveSession(authCookie.Value)
		cookie.RemoveAuthCookie(w)
	}

	http.Redirect(w, r, config.DefaultEndpoints.LoginPage, http.StatusSeeOther)
}
