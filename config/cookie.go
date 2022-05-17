package config

import (
	"net/http"
	"time"
)

const AuthCookieName string = "auth_session_cookie"
const AuthCookiePath string = "/"

func SetAuthCookie(w http.ResponseWriter, token string, expiryTime time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     AuthCookieName,
		Value:    token,
		Path:     AuthCookiePath,
		Expires:  expiryTime,
		HttpOnly: true,
	})
}

func RemoveAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     AuthCookieName,
		Value:    "",
		Path:     AuthCookiePath,
		Expires:  time.Now(),
		HttpOnly: true,
	})
}

func GetAuthCookie(r *http.Request) (*http.Cookie, bool) {
	cookie, noTokenErr := r.Cookie(AuthCookieName)
	if noTokenErr != nil {
		return &http.Cookie{}, false
	}

	return cookie, true
}
