package middlewares

import (
	"context"
	"net/http"
	"printer/persistence/model"
)

// RequireAuth - builds middelware that
// prevents non-authenticated users from making requests to `next` handler
// protected by this middleware
func BuildRequireAuth(
	GetSessionIfValid func(w http.ResponseWriter, r *http.Request) (*model.Session, bool),
	urlForNotAuthneticated string,
) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			session, isValid := GetSessionIfValid(w, r)
			if isValid {
				// add session to the context for furhter handlers
				r = r.WithContext(context.WithValue(r.Context(), ContextSessionKey, session))
				next(w, r)
			} else {
				http.Redirect(w, r, urlForNotAuthneticated, http.StatusSeeOther)
			}
		}
	}
}
