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

// SetUsersPagesPerMonth - sets number of pages a user can print per month
func (controller userController) SetUsersPagesPerMonth(w http.ResponseWriter, r *http.Request) {
	// retrieveing userID param from the url
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// retrieving the number of pages from request body
	var body struct {
		NumberOfPages int
	}

	json.NewDecoder(r.Body).Decode(&body)

	err = controller.db.SetPagesPerMonth(uint(userID), body.NumberOfPages)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

// AllowUsingPrinter - allows specified user to user printer,
// provided his page limit per month is not exceeded
func (controller userController) AllowUsingPrinter(w http.ResponseWriter, r *http.Request) {
	// retrieveing id param from the url
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.db.AllowUsingPrinter(uint(userID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

// ForbidUsingPrinter - allows specified user to user printer,
// provided his page limit per month is not exceeded
func (controller userController) ForbidUsingPrinter(w http.ResponseWriter, r *http.Request) {
	// retrieveing id param from the url
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.db.ForbidUsingPrinter(uint(userID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
