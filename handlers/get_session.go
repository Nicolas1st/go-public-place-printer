package auth

import (
	"net/http"
	"printer/config"
	"printer/persistence/model"
	"printer/pkg/cookie"
)

type sessioner interface {
	GetSessionByToken(sessionToken string) (*model.Session, error)
}

func GetSession(sessioner sessioner, w http.ResponseWriter, r *http.Request) (session *model.Session, doRedirect bool) {
	authCookie, ok := cookie.GetAuthCookie(r)
	// the user is not authenticated
	if !ok {
		http.Redirect(w, r, config.DefaultEndpoints.LoginPage, http.StatusSeeOther)
		return &model.Session{}, true
	}

	session, err := sessioner.GetSessionByToken(authCookie.Value)
	if err != nil {
		http.Redirect(w, r, config.DefaultEndpoints.LoginPage, http.StatusSeeOther)
		return &model.Session{}, true
	}

	return session, false
}
