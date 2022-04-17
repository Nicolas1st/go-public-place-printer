package views

import (
	"fmt"
	"net/http"
	"printer/interfaces"
)

func BuildSignup(templates interfaces.Templates) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			signinTemplate := templates.GetSignup()

			// data query logic
			// to be added in the future
			// mock data
			data := struct {
				Greeting string
			}{
				Greeting: "Hello",
			}

			err := signinTemplate.Execute(w, data)
			fmt.Println(err)
		},
	)
}
