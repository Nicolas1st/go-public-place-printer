package model

import "time"

const ExpiryPeriod time.Duration = 5 * time.Minute

type Session struct {
	UserID    uint
	Username  string
	ExpiresAt time.Time
}

func NewSession(userID uint, username string) *Session {
	return &Session{
		UserID:    userID,
		Username:  username,
		ExpiresAt: time.Now().Add(ExpiryPeriod),
	}
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
