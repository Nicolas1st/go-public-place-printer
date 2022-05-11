package middlewares

import (
	"fmt"
	"net/http"
	"printer/persistence/model"
)

func BuildOnlyAdmin(urlForNonAdmins string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			session := r.Context().Value(ContextSessionKey).(*model.Session)
			fmt.Println(session.User.Role)
			if session.User.Role == model.Admin {
				next(w, r)
			} else {
				http.Redirect(w, r, urlForNonAdmins, http.StatusSeeOther)
			}
		}
	}
}
