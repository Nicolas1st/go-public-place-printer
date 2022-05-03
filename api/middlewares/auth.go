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
		// check if auth cookie is present
		cookie, err := r.Cookie(auth.AuthCookieName)
		if err != nil {
			redirectTo(w, r)
			return
		}

		// checking whether a valid sesion exists for the provided
		session, err := sessionStorage.GetSessionByToken(cookie.Value)
		if err != nil {
			redirectTo(w, r)
		}

		// storing the session in the context
		r = r.WithContext(context.WithValue(r.Context(), ContextSessionKey, session))
		next.ServeHTTP(w, r)
	}
}

// OnlyAnonymous - prevents authenticated users from making requests to certain enpoints (parameter next)
// based on whether the auth cookie is present
func OnlyAnonymous(next *http.ServeMux, redirectTo http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie(auth.AuthCookieName)

		if err == nil {
			// non authenticated don't have auth cookie
			redirectTo(w, r)
		} else {
			// authenticated users have one
			next.ServeHTTP(w, r)
		}
	}
}
