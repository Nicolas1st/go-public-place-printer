package model

import "time"

type Session struct {
	UserID    uint
	Username  string
	User      *User
	ExpiresAt time.Time
}

func NewSession(user *User, expiryTime time.Time) *Session {
	return &Session{
		User:      user,
		UserID:    user.ID,
		Username:  user.Name,
		ExpiresAt: expiryTime,
	}
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
