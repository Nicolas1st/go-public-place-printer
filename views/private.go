package views

import (
	"net/http"
	"printer/api/middlewares"
	"printer/api/views/pages"
	"printer/persistence/model"
)

type DatabaseInterface interface {
	GetAllPrintsByUID(UID uint) []model.Print
}

func BuildSubmitFileView(page *pages.SubmitFilePage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := page.Execute(w)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func BuildProfileView(page *pages.ProfilePage, database DatabaseInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(middlewares.ContextSessionKey).(*model.Session)
		user := session.User

		data := pages.ProfilePageData{
			User:   *user,
			Prints: database.GetAllPrintsByUID(user.ID),
		}

		err := page.Execute(w, data)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
