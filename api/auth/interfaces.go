package auth

import (
	"printer/persistence/db"
	"printer/persistence/model"
	"time"
)

type SessionStorageInterface interface {
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
}

type DatabaseInterface interface {
	GetUserByName(username string) (*db.User, error)
}
