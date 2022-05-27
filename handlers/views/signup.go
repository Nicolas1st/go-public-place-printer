package views

import (
	"net/http"
	"printer/handlers"
	"printer/handlers/views/pages"
)

func (c *viewsController) buildSignUpView(p *pages.Pages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if messages, success := c.handleSignUpForm(w, r); !success {
				p.Public.SignUp.Execute(w, messages, nil)
			} else {
				http.Redirect(w, r, handlers.DefaultEndpoints.LoginPage, http.StatusSeeOther)
			}

			return
		}

		p.Public.SignUp.Execute(w, pages.NewFlashMessages(), nil)
	}
}

func (c *viewsController) handleSignUpForm(w http.ResponseWriter, r *http.Request) (*pages.FlashMessages, bool) {
	// достать данные из формы
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	repeatPassword := r.PostFormValue("repeatPassword")

	messages := pages.NewFlashMessages()

	// проверить уникальность логина
	if _, err := c.db.GetUserByName(username); err == nil {
		messages.Add("Логин уже занят", pages.ErrorMessage)
	}

	// проверить уникальность email
	if _, err := c.db.GetUserByEmail(email); err == nil {
		messages.Add("Email уже занят", pages.ErrorMessage)
	}

	// проверить длину логина
	if len(username) < 8 || len(username) > 20 {
		messages.Add("Логин должен иметь длину от 8 до 20 символов", pages.ErrorMessage)
	}

	// проверить совпадение паролей
	if password != repeatPassword {
		messages.Add("Пароли не совадают", pages.ErrorMessage)
	} else if len(password) < 8 {
		messages.Add("Пароль должен состоять из более чем 8 символов", pages.ErrorMessage)
	}

	if !messages.HasErrorMessages() {
		// создать новый аккаунт
		err := c.db.CreateNewUser(username, email, password)
		return messages, err == nil
	}

	return messages, false
}
