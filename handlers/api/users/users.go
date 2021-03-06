package users

import (
	"net/http"
	"printer/handlers"
	"printer/persistence/db"

	"github.com/gorilla/mux"
)

type userController struct {
	db databaseInterface
}

func NewApi(db *db.Database) *mux.Router {
	// api dependencies
	controller := userController{db: db}

	// api router
	r := mux.NewRouter()

	// users
	r.HandleFunc(handlers.DefaultEndpoints.UsersApi+"{id:[0-9]+}", controller.GetUser).Methods(http.MethodGet)
	r.HandleFunc(handlers.DefaultEndpoints.UsersApi, controller.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc(handlers.DefaultEndpoints.UsersApi+"{id:[0-9]+}", controller.DeleteUser).Methods(http.MethodDelete)

	// permissions
	r.HandleFunc(handlers.DefaultEndpoints.UsersApi+"{id:[0-9]+}/printing/permission", controller.AllowUsingPrinter).Methods(http.MethodPatch)
	r.HandleFunc(handlers.DefaultEndpoints.UsersApi+"{id:[0-9]+}/printing/prohibition", controller.ForbidUsingPrinter).Methods(http.MethodPatch)
	r.HandleFunc(handlers.DefaultEndpoints.UsersApi+"{id:[0-9]+}/pages", controller.SetUsersPagesPerMonth).Methods(http.MethodPatch)

	return r
}
