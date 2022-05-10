package register

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	DB DatabaseInterface
}

func (controller UserController) BuildCreateNewUser(
	redirectToOnSuccess http.HandlerFunc,
	redirectToOnFail http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		repeatPassword := r.PostFormValue("repeatPassword")

		// making sure the password are the same on the server side
		if password != repeatPassword {
			redirectToOnFail(w, r)
			return
		}

		// hashing the password
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
		if err != nil {
			// add flash messages later
			redirectToOnFail(w, r)
			return
		}

		// creating new user
		err = controller.DB.CreateNewUser(username, email, string(passwordHash))

		// abort if it was possible to create the user account
		if err != nil {
			redirectToOnFail(w, r)
			return
		}

		redirectToOnSuccess(w, r)
	}
}
