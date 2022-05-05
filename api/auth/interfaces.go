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
	GetUserByName(username string) (*model.User, error)
}
