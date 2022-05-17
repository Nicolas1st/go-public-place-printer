package views

import (
	"net/http"
	"printer/handlers"
	"printer/handlers/views/pages"
)

func (c *viewsController) buildSignUpView(p *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if messages, success := c.handleSignUpForm(w, r); !success {
				p.Execute(w, messages, nil)
			} else {
				http.Redirect(w, r, handlers.DefaultEndpoints.LoginPage, http.StatusSeeOther)
			}

			return
		}

		p.Execute(w, pages.NewFlashMessages(), nil)
	}
}

func (c *viewsController) handleSignUpForm(w http.ResponseWriter, r *http.Request) (*pages.FlashMessages, bool) {
	// extract data from form
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	repeatPassword := r.PostFormValue("repeatPassword")

	messages := pages.NewFlashMessages()

	// check username uniqness
	if _, err := c.db.GetUserByName(username); err == nil {
		messages.Add("The username provided is already occupied", pages.ErrorMessage)
	}

	// check email uniqness
	if _, err := c.db.GetUserByEmail(email); err != nil {
		messages.Add("The email provided is already occupied", pages.ErrorMessage)
	}

	// check passwords match
	if password != repeatPassword {
		messages.Add("The passwords do not match", pages.ErrorMessage)
	}

	if !messages.HasErrorMessages() {
		// create a new account
		err := c.db.CreateNewUser(username, email, password)
		return messages, err == nil
	}

	return messages, true
}
