package users

import (
	"encoding/json"
	"net/http"
	"printer/persistence/db"
)

type userController struct {
	db *db.Database
}

func (controller userController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var submittedData createUserRequestBody
	json.NewDecoder(r.Body).Decode(&submittedData)

	if err := controller.db.CreateNewUser(submittedData.name, submittedData.passwordHash); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (controller userController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	var submittedData deleteUserByIDRequestBody
	json.NewDecoder(r.Body).Decode(&submittedData)

	err := controller.db.DeleteUserByID(submittedData.UID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (controller userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := controller.db.GetAllUsers()

	json.NewEncoder(w).Encode(users)
}

func (controller userController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	var submittedData getUserByIDRequestBody
	json.NewDecoder(r.Body).Decode(&submittedData)

	user, err := controller.db.GetUserByID(submittedData.UID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
