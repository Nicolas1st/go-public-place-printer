package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `gorm:"unique"`
	Email         string `gorm:"unique"`
	PasswordHash  string
	PagesPerMonth int  `gorm:"default:100"`
	CanUsePrinter bool `gorm:"default:true"`
}
