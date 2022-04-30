package auth

import (
	"printer/persistence/model"
	"time"
)

type SessionStorageInterface interface {
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
}

type DatabaseInterface interface {
	GetUserByUsername(username string) (*model.User, error)
}
