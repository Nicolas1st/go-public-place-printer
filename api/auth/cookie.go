package auth

import (
	"net/http"
	"time"
)

const AuthCookieName string = "auth_session_cookie"
const CookiePath string = "/"

func SetAuthCookie(w http.ResponseWriter, r *http.Request, token string, expiryTime time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     AuthCookieName,
		Value:    token,
		Path:     CookiePath,
		Expires:  expiryTime,
		HttpOnly: true,
	})
}

func RemoveAuthCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     AuthCookieName,
		Value:    "",
		Path:     CookiePath,
		Expires:  time.Now(),
		HttpOnly: true,
	})
}
