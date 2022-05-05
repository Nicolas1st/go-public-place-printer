package users

import (
	"net/http"
	"printer/persistence/db"

	"github.com/gorilla/mux"
)

func NewRouter(db *db.Database) *mux.Router {
	// api dependencies
	controller := userController{db: db}

	// api router
	r := mux.NewRouter()

	r.HandleFunc("/{id:[0-9]+}", controller.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/", controller.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/{id:[0-9]+}", controller.DeleteUser).Methods(http.MethodDelete)

	return r
}
