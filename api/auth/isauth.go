package auth

import (
	"net/http"
	"printer/persistence/model"
)

func GetSessionToken(w http.ResponseWriter, r *http.Request) (sessionToken string, noTokenErr error) {
	// check if user has the auth cookie
	cookie, noTokenErr := r.Cookie(AuthCookieName)
	if noTokenErr != nil {
		return "", noTokenErr
	}

	sessionToken = cookie.Value
	return sessionToken, noTokenErr
}

func (resource *authDependencies) GetSessionIfValid(w http.ResponseWriter, r *http.Request) (session *model.Session, valid bool) {
	// check if user has the auth cookie
	sessionToken, noTokenErr := GetSessionToken(w, r)
	if noTokenErr != nil {
		return &model.Session{}, false
	}

	// check whether there is corresponding session in server's memory
	session, noSessionErr := resource.sessionStorage.GetSessionByToken(sessionToken)
	if noSessionErr != nil {
		// remove cookie if there is no correspoding session on the server
		RemoveAuthCookie(w, r)
		return session, false
	}

	// check for expiration
	if session.IsExpired() {
		// remove cookie if the session is expired
		RemoveAuthCookie(w, r)
		// remove the session in the storage
		resource.sessionStorage.RemoveSession(sessionToken)
		return session, false
	}

	return session, true
}
