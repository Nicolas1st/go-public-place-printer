package session

import (
	"printer/persistence/model"
	"time"

	"github.com/google/uuid"
)

type SessionStorage struct {
	sessions     map[string]*model.Session
	lastPurgedAt time.Time
	purgePeriod  time.Duration
}

func NewSessionStorage(purgePeriod time.Duration) *SessionStorage {
	return &SessionStorage{
		sessions:     make(map[string]*model.Session),
		lastPurgedAt: time.Now(),
		purgePeriod:  purgePeriod,
	}
}

// purgeFromExpiredSessions - purges storage from already expired sessions
// that were not yet removed explicitly
func (s *SessionStorage) purgeFromExpiredSessions() {
	for token, session := range s.sessions {
		if session.IsExpired() {
			delete(s.sessions, token)
		}
	}
}

// StoreSession - stores session in the storage,
// before storing a new one, the function purges
// already expires sessions
func (s *SessionStorage) StoreSession(user *model.User) (string, time.Time) {
	// create session
	session := model.NewSession(user, time.Now().Add(s.purgePeriod))

	// to avoid memory leaks the session are being purged
	// It's done every expiry perdiod of one cookies elapses
	// the persiod is defined in session.go
	if time.Now().After(s.lastPurgedAt.Add(s.purgePeriod)) {
		s.purgeFromExpiredSessions()
	}

	defer func() {
		// the token generator function for some reason can throw a panic
		// it's an inner implementation issue(?), so it's being handled here
		recover()
	}()

	// in case the already used token is being generated,
	// it's almost impossible but can happen anyway
	var sessionToken string
	for {
		sessionToken = uuid.NewString()
		if _, alreadyExists := s.sessions[sessionToken]; !alreadyExists {
			break
		}
	}

	// storing the session in memory
	s.sessions[sessionToken] = session

	return sessionToken, session.ExpiresAt
}

// RemoveSession - removes session associated with the provided token from the storage
func (s *SessionStorage) RemoveSession(sessionToken string) {
	delete(s.sessions, sessionToken)
}

// GetSessionByToken checks whether the session is valid,
// it checks if it exists and is not too old
func (s *SessionStorage) GetSessionByToken(sessionToken string) (*model.Session, bool) {
	session, exists := s.sessions[sessionToken]
	if !exists {
		return &model.Session{}, false
	}

	if session.IsExpired() {
		s.RemoveSession(sessionToken)
		return &model.Session{}, false
	}

	return session, true
}
