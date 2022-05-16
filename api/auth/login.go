package auth

import (
	"errors"
	"net/http"
	"printer/persistence/model"

	"golang.org/x/crypto/bcrypt"
)

func (c *authController) Authenticate(w http.ResponseWriter, r *http.Request) error {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// check if user exists
	user, err := c.db.GetUserByName(username)
	if err != nil {
		return errors.New("not user found with the name specified")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return errors.New("username or password is wrong")
	}

	// create session
	session := model.NewSession(user, user.Name)
	token, expiryTime := c.sessions.StoreSession(session)

	// set session cookie in the user's browser
	setAuthCookie(w, token, expiryTime)

	return nil
}
