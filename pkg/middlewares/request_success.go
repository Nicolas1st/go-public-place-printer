package middlewares

import "net/http"

func BuildRedirectOnRequest(successURL, failureURL string) ApiRequestMiddleware {
	return func(
		next func(w http.ResponseWriter, r *http.Request) error,
	) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := next(w, r)
			if err == nil {
				http.Redirect(w, r, successURL, http.StatusSeeOther)
			} else {
				http.Redirect(w, r, failureURL, http.StatusSeeOther)
			}
		}
	}
}
