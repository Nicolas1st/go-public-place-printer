package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type userController struct {
	db databaseInterface
}

// GetUser - return one user if id is provided
func (controller userController) GetUser(w http.ResponseWriter, r *http.Request) {
	// retrieveing id param from the url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := controller.db.GetUserByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetAllUsers - returns all users
func (controller userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := controller.db.GetAllUsers()

	json.NewEncoder(w).Encode(users)
}

// DeleteUser - deletes one user if id is provided
func (controller userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// retrieveing id param from the url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.db.DeleteUserByID(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
