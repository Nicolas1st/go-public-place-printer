package handlers

import "net/http"

func ForNotLoggedIn(sessioner Sessioner, w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, ok := GetSession(sessioner, w, r)
	if ok {
		http.Redirect(w, r, DefaultEndpoints.PrinterPage, http.StatusSeeOther)
	} else {
		next(w, r)
	}
}

func ForCommonUsers(sessioner Sessioner, w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, ok := GetSession(sessioner, w, r)
	if !ok {
		http.Redirect(w, r, DefaultEndpoints.LoginPage, http.StatusSeeOther)
	} else {
		next(w, r)
	}
}

func ForAdmin(sessioner Sessioner, w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session, ok := GetSession(sessioner, w, r)
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
