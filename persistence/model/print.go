package model

import "gorm.io/gorm"

type Print struct {
	gorm.Model
	Filename      string
	Filepath      string
	NumberOfPages int
	Username      string
	UserID        uint
	User          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
