package account

import (
	"encoding/json"
	"net/http"
)

// Request json schema contents the handler expects to receive
type createAccountRequestSchema struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

// Response json schema contents sent back by the handler
type createAccountResponseSchema struct {
	RedirectionURL string   `json:"redirectionURL"`
	ErrorMessages  []string `json:"flashMessage"`
}

func (c *accountController) validateNewAccountData(schema createAccountRequestSchema) (errorMessages []string, valid bool) {
	valid = true

	// check username uniqness
	if _, err := c.db.GetUserByName(schema.Username); err == nil {
		errorMessages = append(errorMessages, "The username provided is already occupied")
		valid = false
	}

	// check email uniqness
	if _, err := c.db.GetUserByEmail(schema.Email); err != nil {
		errorMessages = append(errorMessages, "The email provided is already occupied")
		valid = false
	}

	// check passwords match
	if schema.Password != schema.RepeatPassword {
		errorMessages = append(errorMessages, "The passwords do not match")
		valid = false
	}

	return
}

// CreateAccount - creates an accountCreateAccountResponse,
// if creation is successful, provides a redirectin link
// otherwise provides a list of error messages to be displayed
// on the sign up page
func (c *accountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var requestFields createAccountRequestSchema

	// decode request, if not possible send bad request response status
	err := json.NewDecoder(r.Body).Decode(&requestFields)
	if err != nil {
		http.Error(w, "Could not parse the request", http.StatusBadRequest)
		return
	}

	var response createAccountResponseSchema

	// check request data validity, if not valid, send back the errors
	errorMessages, valid := c.validateNewAccountData(requestFields)
	response.ErrorMessages = errorMessages

	if !valid {
		json.NewEncoder(w).Encode(response)
		return
	}

	// create a new account
	err = c.db.CreateNewUser(requestFields.Username, requestFields.Email, requestFields.Password)
	if err != nil {
		http.Error(w, "Could not create account", http.StatusInternalServerError)
		return
	}

	response.RedirectionURL = "/signin"
	json.NewEncoder(w).Encode(response)
}
