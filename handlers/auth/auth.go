package auth

import (
	"printer/persistence/model"
	"time"
)

type sessionStorage interface {
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
	GetSessionByToken(sessionToken string) (*model.Session, error)
}

type database interface {
	GetUserByName(username string) (*model.User, error)
}

type authController struct {
	sessions sessionStorage
	db       database
}

func NewController(sessions sessionStorage, db database) *authController {
	return &authController{
		sessions: sessions,
		db:       db,
	}
}
