package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AllowUsingPrinterResponse struct {
	Permission bool `json:"permission"`
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

	permission, err := controller.db.AllowUsingPrinter(uint(userID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		jsonResponse := AllowUsingPrinterResponse{Permission: permission}
		json.NewEncoder(w).Encode(&jsonResponse)
	}
}
