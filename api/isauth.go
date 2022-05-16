package auth

import (
	"net/http"
	"printer/config"
	"printer/persistence/model"
)

func getSessionToken(w http.ResponseWriter, r *http.Request) (sessionToken string, noTokenErr error) {
	// check if user has the auth cookie
	cookie, noTokenErr := r.Cookie(config.AuthCookieName)
	if noTokenErr != nil {
		return "", noTokenErr
	}

	sessionToken = cookie.Value
	return sessionToken, noTokenErr
}

func (resource *authController) GetSessionIfValid(w http.ResponseWriter, r *http.Request) (session *model.Session, valid bool) {
	// check if user has the auth cookie
	sessionToken, noTokenErr := getSessionToken(w, r)
	if noTokenErr != nil {
		return &model.Session{}, false
	}

	// check whether there is corresponding session in server's memory
	session, noSessionErr := resource.sessions.GetSessionByToken(sessionToken)
	if noSessionErr != nil {
		// remove cookie if there is no correspoding session on the server
		removeAuthCookie(w)
		return session, false
	}

	// check for expiration
	if session.IsExpired() {
		// remove cookie if the session is expired
		removeAuthCookie(w)
		// remove the session in the storage
		resource.sessions.RemoveSession(sessionToken)
		return session, false
	}

	return session, true
}
