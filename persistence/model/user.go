package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role int

const (
	Admin Role = iota
	NonAdmin
)

type User struct {
	gorm.Model
	Name          string `gorm:"unique;size:128;not null"`
	Email         string `gorm:"unique;size:128;not null"`
	Role          Role   `gorm:"not null"`
	PasswordHash  string `gorm:"not null"`
	PagesPerMonth uint   `gorm:"default:100"`
	CanUsePrinter bool   `gorm:"default:true"`
}

func (u *User) IsPasswordValid(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (u *User) IsAdmin() bool {
	return u.Role == Admin
}
