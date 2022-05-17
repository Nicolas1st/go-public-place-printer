package middlewares

import (
	"net/http"
	"printer/persistence/model"
)

// BuildOnlyNotAuth - builds middleware that
// prevents authenticated users from making requests to certain enpoints (parameter next)
// based on whether the auth cookie is present
func BuildRequireNotAuth(
	GetSessionIfValid func(w http.ResponseWriter, r *http.Request) (*model.Session, bool),
	urlForAuthenticated string,
) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_, isSessionValid := GetSessionIfValid(w, r)
			if !isSessionValid {
				next(w, r)
			} else {
				http.Redirect(w, r, urlForAuthenticated, http.StatusSeeOther)
			}
		}
	}
}
