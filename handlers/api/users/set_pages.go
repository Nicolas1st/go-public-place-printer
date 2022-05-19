package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
