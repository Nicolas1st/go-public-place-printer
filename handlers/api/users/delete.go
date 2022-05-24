package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DeleteUserResponse struct {
	Result string `json:"result"`
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
	} else {
		jsonResponse := DeleteUserResponse{Result: "deleted"}
		json.NewEncoder(w).Encode(&jsonResponse)
	}
}
