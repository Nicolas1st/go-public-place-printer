package auth

import (
	"net/http"
	"printer/config"
)

func (c *authController) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(config.AuthCookieName)
	if err != http.ErrNoCookie {
		// remove session on the server
		c.sessions.RemoveSession(cookie.Value)

		// remove session cookie in the browser
		removeAuthCookie(w)
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
