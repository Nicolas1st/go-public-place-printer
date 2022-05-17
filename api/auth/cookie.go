package auth

import (
	"net/http"
	"printer/config"
	"time"
)

func setAuthCookie(w http.ResponseWriter, token string, expiryTime time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     config.AuthCookieName,
		Value:    token,
		Path:     config.AuthCookiePath,
		Expires:  expiryTime,
		HttpOnly: true,
	})
}

func removeAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     config.AuthCookieName,
		Value:    "",
		Path:     config.AuthCookiePath,
		Expires:  time.Now(),
		HttpOnly: true,
	})
}
