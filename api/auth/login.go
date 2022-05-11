package auth

import (
	"errors"
	"net/http"
	"printer/persistence/model"

	"golang.org/x/crypto/bcrypt"
)

func (resource *authDependencies) Authenticate(w http.ResponseWriter, r *http.Request) error {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// check if user exists
	user, err := resource.database.GetUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("not user found with the name specified")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return errors.New("username or password is wrong")
	}

	// create session
	session := model.NewSession(user.ID, user.Name)

	// store session in memory
	token, expiryTime := resource.sessionStorage.StoreSession(session)

	// set session cookie in the user's browser
	SetAuthCookie(w, r, token, expiryTime)

	return nil
}
