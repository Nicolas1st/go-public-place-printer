package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SetPagesRequest struct {
	NumberOfPages uint `json:"numberOfPages"`
}

type SetPagesResponse struct {
	NumberOfPages uint `json:"numberOfPages"`
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

	// retrieving number of pages from request body
	var jsonRequest SetPagesRequest
	json.NewDecoder(r.Body).Decode(&jsonRequest)

	pagesPerMonth, err := controller.db.SetPagesPerMonth(uint(userID), jsonRequest.NumberOfPages)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		jsonResponse := SetPagesResponse{NumberOfPages: pagesPerMonth}
		json.NewEncoder(w).Encode(&jsonResponse)
	}
}
