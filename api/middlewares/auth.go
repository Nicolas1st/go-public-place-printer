package middlewares

import (
	"context"
	"net/http"
	"printer/api/auth"
	"printer/persistence/model"
)

type contextSessionKey string

var ContextSessionKey contextSessionKey = "contextSessionKey"

type SessionStorageInterface interface {
	GetSessionByToken(sessionToken string) (*model.Session, error)
}

// OnlyAuthenticated - prevents non-authenticated users from making requests to certain enpoints (parameter next)
// based on whether the auth cookie is present
// and the validity of the session
func OnlyAuthenticated(sessionStorage SessionStorageInterface, next *http.ServeMux, redirectTo http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(auth.AuthCookieName)

		// check if auth cookie is present
		if err != nil {
			redirectTo(w, r)
			return
		}

		// extracing the session
		session, err := sessionStorage.GetSessionByToken(cookie.Value)

		// in case of error the session is considered to be non-existsent
		if err != nil {
			// storing the session in the context
			r = r.WithContext(context.WithValue(r.Context(), ContextSessionKey, session))

			next.ServeHTTP(w, r)
			return
		}

		// redirect to another view if user is not authenticated
		redirectTo(w, r)
	}
}

// OnlyAnonymous - prevents authenticated users from making requests to certain enpoints (parameter next)
// based on whether the auth cookie is present
func OnlyAnonymous(next *http.ServeMux, redirectTo http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, cookieNotPresentErr := r.Cookie(auth.AuthCookieName)

		// redirect to another view if cookie is present
		if cookieNotPresentErr == nil {
			redirectTo(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
