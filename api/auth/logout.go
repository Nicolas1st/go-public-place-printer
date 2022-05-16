package auth

import (
	"net/http"
	"printer/config"
)

func (c *authController) Logout(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(config.AuthCookieName)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	// remove session on the server
	c.sessions.RemoveSession(cookie.Value)

	// remove session cookie in the browser
	removeAuthCookie(w)

	return nil
}
