package users

import "printer/persistence/model"

type databaseInterface interface {
	// user managment
	DeleteUserByID(id uint) error
	GetAllUsers() []model.User
	GetUserByID(id uint) (*model.User, error)

	// permissions
	SetPagesPerMonth(userID uint, pagesPerMonth uint) (uint, error)
	AllowUsingPrinter(userID uint) (bool, error)
	ForbidUsingPrinter(userID uint) (bool, error)
}
