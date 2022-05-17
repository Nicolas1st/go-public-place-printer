package users

import "printer/persistence/model"

type databaseInterface interface {
	// user managment
	DeleteUserByID(id uint) error
	GetAllUsers() []model.User
	GetUserByID(id uint) (*model.User, error)

	// permissions
	SetPagesPerMonth(id uint, pagesPerMonth int) error
	AllowUsingPrinter(userID uint) error
	ForbidUsingPrinter(userID uint) error
}
