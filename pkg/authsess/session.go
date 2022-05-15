package model

import "time"

const ExpiryPeriod time.Duration = 5 * time.Minute

type Session struct {
	UserID    uint
	Username  string
	User      *User
	ExpiresAt time.Time
}

func NewSession(user *User, username string) *Session {
	return &Session{
		User:      user,
		UserID:    user.ID,
		Username:  username,
		ExpiresAt: time.Now().Add(ExpiryPeriod),
	}
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
