package auth

import (
	"net/http"
	"time"
)

const authCookieName string = "auth_session_cookie"
const cookiePath string = "/"

func setAuthCookie(w http.ResponseWriter, token string, expiryTime time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     authCookieName,
		Value:    token,
		Path:     cookiePath,
		Expires:  expiryTime,
		HttpOnly: true,
	})
}

func removeAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     authCookieName,
		Value:    "",
		Path:     cookiePath,
		Expires:  time.Now(),
		HttpOnly: true,
	})
}
