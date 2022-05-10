package model

import "gorm.io/gorm"

type Role int

const (
	Admin Role = iota
	NonAdmin
)

type User struct {
	gorm.Model
	Name          string `gorm:"unique"`
	Email         string `gorm:"unique"`
	Role          Role
	PasswordHash  string
	PagesPerMonth int  `gorm:"default:100"`
	CanUsePrinter bool `gorm:"default:true"`
}
