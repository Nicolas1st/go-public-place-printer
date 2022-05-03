package auth

import (
	"net/http"
	"printer/persistence/model"

	"golang.org/x/crypto/bcrypt"
)

func (resource *AuthResource) authenticate(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// check if user exists
	user, err := resource.database.GetUserByName(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create session
	session := model.NewSession(user.ID, user.Name)

	// store session in memory
	token, expiryTime := resource.sessionStorage.StoreSession(session)

	// set session cookie in the user's browser
	http.SetCookie(w, &http.Cookie{
		Name:     AuthCookieName,
		Value:    token,
		Path:     CookiePath,
		Expires:  expiryTime,
		HttpOnly: true,
	})
}
