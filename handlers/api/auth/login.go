package auth

import (
	"encoding/json"
	"net/http"
	"printer/persistence/model"
	"printer/pkg/cookie"
)

type loginRequestSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponseSchema struct {
	RedirectionURL string   `json:"redirectionURL"`
	ErrorMessages  []string `json:"flashMessage"`
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var jsonRequest loginRequestSchema
	if err := json.NewDecoder(r.Body).Decode(&jsonRequest); err != nil {
		http.Error(w, "Could not parse the request", http.StatusBadRequest)
		return
	}

	var jsonResponse loginResponseSchema

	// check if user exists
	user, err := c.db.GetUserByName(jsonRequest.Username)
	if err != nil {
		jsonResponse.ErrorMessages = append(jsonResponse.ErrorMessages, "Incorrect name or password", jsonRequest.Username)
		json.NewEncoder(w).Encode(jsonResponse)
		return
	}

	// check if password is valid
	if !user.IsPasswordValid(jsonRequest.Password) {
		jsonResponse.ErrorMessages = append(jsonResponse.ErrorMessages, "Incorrect name or password", jsonRequest.Username)
		json.NewEncoder(w).Encode(jsonResponse)
		return
	}

	// create session
	session := model.NewSession(user, user.Name)
	token, expiryTime := c.sessions.StoreSession(session)

	// set session cookie in the user's browser
	cookie.SetAuthCookie(w, token, expiryTime)

	// set redirection url
	jsonResponse.RedirectionURL = "/submit-file"

	json.NewEncoder(w).Encode(jsonResponse)
}
