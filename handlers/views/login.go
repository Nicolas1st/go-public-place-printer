package views

import (
	"net/http"
	"printer/handlers"
	"printer/handlers/views/pages"
	"printer/persistence/model"
)

func (c *viewsController) buildLoginView(p *pages.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if messages, success := c.handleLoginForm(w, r); !success {
				p.Execute(w, messages, nil)
			} else {
				http.Redirect(w, r, handlers.DefaultEndpoints.PrinterPage, http.StatusSeeOther)
			}

			return
		}

		p.Execute(w, pages.NewFlashMessages(), nil)
	}
}

func (c *viewsController) handleLoginForm(w http.ResponseWriter, r *http.Request) (*pages.FlashMessages, bool) {
	// extract data from form
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	messages := pages.NewFlashMessages()

	// check if user exists
	user, err := c.db.GetUserByName(username)
	if err != nil {
		messages.Add("Incorrect name or password", pages.ErrorMessage)
		return messages, false
	}

	// check if password is valid
	if !user.IsPasswordValid(password) {
		messages.Add("Incorrect name or password", pages.ErrorMessage)
		return messages, false
	}

	// create session
	session := model.NewSession(user, user.Name)
	token, expiryTime := c.sessioner.StoreSession(session)

	// set session cookie in the user's browser
	handlers.SetAuthCookie(w, token, expiryTime)

	return messages, true
}
