package model

import "time"

const ExpiryPeriod time.Duration = 5 * time.Minute

type Session struct {
	UserID    uint
	ExpiresAt time.Time
}

func NewSession(userID uint) *Session {
	return &Session{
		UserID:    userID,
		ExpiresAt: time.Now().Add(ExpiryPeriod),
	}
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
