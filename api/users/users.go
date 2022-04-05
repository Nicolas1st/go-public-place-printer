package users

import (
	"net/http"
	"printer/persistence/db"

	"github.com/gorilla/mux"
)

func NewRouter(db *db.Database) *mux.Router {
	r := mux.NewRouter()

	controller := userController{}

	r.HandleFunc("/", controller.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/", controller.GetAllUsers).Methods(http.MethodGet)

	r.HandleFunc("/", controller.DeleteUserByID).Methods(http.MethodDelete)
	r.HandleFunc("/", controller.GetAllUsers).Methods(http.MethodGet)

	return r
}
