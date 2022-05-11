package auth

import (
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
