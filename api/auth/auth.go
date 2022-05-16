package auth

import (
	"net/http"
	"printer/persistence/model"
	"time"
)

type SessionStorageInterface interface {
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
	GetSessionByToken(sessionToken string) (*model.Session, error)
}

type DatabaseInterface interface {
	GetUserByName(username string) (*model.User, error)
}

type authDependencies struct {
	sessionStorage SessionStorageInterface
	database       DatabaseInterface
}

type AuthHandlers struct {
	Login             func(w http.ResponseWriter, r *http.Request) error
	Logout            func(w http.ResponseWriter, r *http.Request) error
	GetSessionIfValid func(w http.ResponseWriter, r *http.Request) (*model.Session, bool)
}

func NewAuthHandlers(sessionStorage SessionStorageInterface, database DatabaseInterface) *AuthHandlers {
	authDependencies := &authDependencies{
		sessionStorage: sessionStorage,
		database:       database,
	}

	return &AuthHandlers{
		Login:             authDependencies.Authenticate,
		Logout:            authDependencies.Logout,
		GetSessionIfValid: authDependencies.GetSessionIfValid,
	}
}
