package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `gorm:"unique"`
	Email         string `gorm:"unique"`
	PasswordHash  string
	PagesPerMonth int  `gorm:"default:100"`
	CanUsePrinter bool `gorm:"default:true"`
}

type Print struct {
	gorm.Model
	SubmittedFileName string
	StoredFileName    string `gorm:"unique"`
	NumberOfPages     int
	UserID            uint
	User              User
}
