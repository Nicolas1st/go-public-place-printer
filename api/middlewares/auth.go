package middlewares

import (
	"fmt"
	"net/http"
	"printer/api/auth"
)

type SessionStorageInterface interface {
	IsSessionValid(sessionToken string) bool
}

// ForbidForAuthenticated - prevents non-authenticated users from making requests to certain enpoints (parameter next)
// based on whether the auth cookie is present
// and the validity of the session
func ForbidForNonAuthenticated(sessionStorage SessionStorageInterface, next *http.ServeMux, redirectTo http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("hit forbid non auth")
		cookie, err := r.Cookie(auth.AuthCookieName)

		// check if auth cookie is present
		if err != nil {
			redirectTo(w, r)
			return
		}

		// checking whether the session for the cookie provied is valid
		if sessionStorage.IsSessionValid(cookie.Value) {
			next.ServeHTTP(w, r)
			return
		}

		// redirect to another view if user is not authenticated
		redirectTo(w, r)
	}
}

// ForbidForAuthenticated - prevents authenticated users from making requests to certain enpoints (parameter next)
// based on whether the auth cookie is present
func ForbidForAuthenticated(next *http.ServeMux, redirectTo http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("hit forbid auth")
		_, cookieNotPresentErr := r.Cookie(auth.AuthCookieName)
		fmt.Print(cookieNotPresentErr)

		// redirect to another view if cookie is present
		if cookieNotPresentErr == nil {
			redirectTo(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
