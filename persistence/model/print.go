package model

import "gorm.io/gorm"

type Print struct {
	gorm.Model
	Filename      string
	NumberOfPages int
	UserID        uint
	User          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
