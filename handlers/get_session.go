package handlers

import (
	"net/http"
	"printer/persistence/model"
)

type Sessioner interface {
	GetSessionByToken(sessionToken string) (*model.Session, bool)
}

func GetSession(sessioner Sessioner, r *http.Request) (session *model.Session, doRedirect bool) {
	authCookie, ok := GetAuthCookie(r)
	// the user is not authenticated
	if !ok {
		return &model.Session{}, true
	}

	// the token has expired
	session, ok = sessioner.GetSessionByToken(authCookie.Value)
	if ok {
		return &model.Session{}, true
	}

	return session, false
}
