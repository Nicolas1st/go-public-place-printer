package handlers

import "net/http"

func ForNotLoggedIn(sessioner Sessioner, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := GetSession(sessioner, r)
		if ok {
			http.Redirect(w, r, DefaultEndpoints.PrinterPage, http.StatusSeeOther)
		} else {
			next(w, r)
		}

	}
}

func ForCommonUsers(sessioner Sessioner, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := GetSession(sessioner, r)
		if ok {
			next(w, r)
		} else {
			http.Redirect(w, r, DefaultEndpoints.LoginPage, http.StatusSeeOther)
		}
	}
}

func ForAdmin(sessioner Sessioner, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, ok := GetSession(sessioner, r)
		if !ok {
			http.Redirect(w, r, DefaultEndpoints.LoginPage, http.StatusSeeOther)
			return
		}

		if !session.User.IsAdmin() {
			http.Redirect(w, r, DefaultEndpoints.PrinterPage, http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}
