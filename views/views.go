package views

import (
	"net/http"
	"printer/interfaces"

	"github.com/gorilla/mux"
)

func NewRouter(templates interfaces.Templates) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/login",
		BuildLogin(templates),
	).Methods(http.MethodGet)

	r.Handle("/signup",
		BuildSignup(templates),
	).Methods(http.MethodGet)

	return r
}
