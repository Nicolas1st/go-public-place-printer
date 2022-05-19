package users

import (
	"encoding/json"
	"net/http"
)

// GetAllUsers - returns all users
func (controller userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := controller.db.GetAllUsers()

	json.NewEncoder(w).Encode(users)
}
