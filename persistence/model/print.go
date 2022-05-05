package model

import "gorm.io/gorm"

type Print struct {
	gorm.Model
	SubmittedFileName string
	StoredFileName    string `gorm:"unique"`
	NumberOfPages     int
	UserID            uint
	User              User
}
