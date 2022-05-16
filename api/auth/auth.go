package auth

import (
	"net/http"
	"printer/persistence/model"
	"time"
)

type SessionStorage interface {
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
	GetSessionByToken(sessionToken string) (*model.Session, error)
}

type Database interface {
	GetUserByName(username string) (*model.User, error)
}

type authController struct {
	sessions SessionStorage
	db       Database
}

type AuthHandlers struct {
	Login             func(w http.ResponseWriter, r *http.Request) error
	Logout            func(w http.ResponseWriter, r *http.Request) error
	GetSessionIfValid func(w http.ResponseWriter, r *http.Request) (*model.Session, bool)
}

func NewController(sessions SessionStorage, db Database) *authController {
	return &authController{
		sessions: sessions,
		db:       db,
	}
}
