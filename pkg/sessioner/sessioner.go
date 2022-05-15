package sessioner

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Sessioner interface {
	StoreSession(UID) (SessionToken, time.Time)
	RemoveSession(SessionToken)
	GetSessionByToken(SessionToken) (Session, error)
}

// in-memory session storage
type sessioner struct {
	sessions     map[SessionToken]Session
	lastPurgedAt time.Time
	expiryPeriod time.Duration
}

// New - creates new object conforming to the SessionStorage interface
func New(expiryPeriod time.Duration) Sessioner {
	return &sessioner{
		sessions:     make(map[SessionToken]Session),
		lastPurgedAt: time.Now(),
		expiryPeriod: expiryPeriod,
	}
}

// StoreSession - stores a session in memory
// removes already expired ones if the it's the time
func (storage *sessioner) StoreSession(uid UID) (SessionToken, time.Time) {
	// create session
	session := Session{
		UID:        uid,
		ExpiryTime: time.Now().Add(storage.expiryPeriod),
	}

	// to avoid memory leaks the session are being purged
	if time.Now().After(storage.lastPurgedAt.Add(storage.expiryPeriod)) {
		storage.purgeFromExpiredSessions()
	}

	// to recover from a panic that can be thrown by uuid.NewString
	defer func() {
		recover()
	}()

	// create a new until it's unique
	var sessionToken SessionToken
	for {
		sessionToken = SessionToken(uuid.NewString())
		if _, alreadyExists := storage.sessions[sessionToken]; !alreadyExists {
			break
		}
	}

	storage.sessions[sessionToken] = session

	return sessionToken, session.ExpiryTime.Add(storage.expiryPeriod)
}

// RemoveSession - removes session associated with the provided token from the storage
func (storage *sessioner) RemoveSession(token SessionToken) {
	delete(storage.sessions, token)
}

// GetSessionByToken checks whether the session is valid,
// it checks if it exists and is not too old
func (storage *sessioner) GetSessionByToken(token SessionToken) (Session, error) {
	session, exists := storage.sessions[token]
	if !exists {
		return Session{}, errors.New("session does not exist")
	}

	return session, nil
}

// purgeFromExpiredSessions - purges storage from already expired sessions
// that were not yet removed explicitly
func (storage *sessioner) purgeFromExpiredSessions() {
	for token, session := range storage.sessions {
		if session.ExpiryTime.Add(storage.expiryPeriod).After(time.Now()) {
			delete(storage.sessions, token)
		}
	}
}
