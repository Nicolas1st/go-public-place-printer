package users

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
