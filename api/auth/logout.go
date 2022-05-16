package auth

import "net/http"

func (resource *authController) Logout(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(AuthCookieName)
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	// removing the session by token
	resource.sessions.RemoveSession(cookie.Value)

	// removing the session in the browser
	RemoveAuthCookie(w)

	return nil
}
