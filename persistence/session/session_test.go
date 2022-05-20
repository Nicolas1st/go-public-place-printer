package session

import (
	"printer/persistence/model"
	"testing"
	"time"
)

func Init() {}

func TestSessionStore(t *testing.T) {
	expiryPeriod := 1 * time.Second
	s := NewSessionStorage(expiryPeriod)

	// store
	u := &model.User{}
	token, _ := s.StoreSession(u)

	// check if stored
	if _, ok := s.sessions[token]; !ok {
		t.Errorf("The session was not stored")
	}
}

func TestPrugeFromExpiredSessions(t *testing.T) {
	// the purgeFromExpiredSessions gets called on getSession
	expiryPeriod := 1 * time.Second
	s := NewSessionStorage(expiryPeriod)

	// store
	u := &model.User{}
	token, _ := s.StoreSession(u)

	// check if stored
	s.GetSessionByToken(token)
	if _, ok := s.sessions[token]; !ok {
		t.Errorf("The session was not stored")
	}

	// check if removed after the expiry period
	time.Sleep(expiryPeriod)
	s.GetSessionByToken(token)
	if _, ok := s.sessions[token]; ok {
		t.Errorf("The session was not removed after the specified time")
	}

}

func TestRemoveSession(t *testing.T) {
	expiryPeriod := 1 * time.Second
	s := NewSessionStorage(expiryPeriod)

	// store
	user := &model.User{}
	token, _ := s.StoreSession(user)

	// check if stored
	s.GetSessionByToken(token)
	if _, ok := s.sessions[token]; !ok {
		t.Errorf("The session was not stored")
	}

	// remove
	s.RemoveSession(token)

	// check if was removed
	if _, ok := s.sessions[token]; ok {
		t.Errorf("The session was not removed")
	}
}

func TestGetSessionByToken(t *testing.T) {
	expiryPeriod := 1 * time.Second
	s := NewSessionStorage(expiryPeriod)

	// test get session that does not exist
	if _, ok := s.GetSessionByToken("does not exist"); ok {
		t.Errorf("Returned non existing session")
	}

	// store session
	u := &model.User{}
	token, _ := s.StoreSession(u)

	// check get
	if _, ok := s.GetSessionByToken(token); !ok {
		t.Errorf("Could not get existing session")
	}
}
