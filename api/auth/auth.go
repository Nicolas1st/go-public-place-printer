package auth

import (
	"net/http"
	"printer/persistence/model"
)

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
