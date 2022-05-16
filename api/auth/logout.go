package auth

import (
	"net/http"
	"printer/config"
)

func (resource *authController) Logout(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(config.AuthCookieName)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	// removing the session by token
	resource.sessions.RemoveSession(cookie.Value)

	// removing the session in the browser
	removeAuthCookie(w)

	return nil
}
