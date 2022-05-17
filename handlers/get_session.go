package handlers

import (
	"net/http"
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
		return &model.Session{}, true
	}

	// the token has expired
	session, err := sessioner.GetSessionByToken(authCookie.Value)
	if err != nil {
		return &model.Session{}, true
	}

	return session, false
}
