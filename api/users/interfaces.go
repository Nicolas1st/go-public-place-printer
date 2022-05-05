package users

import "printer/persistence/model"

type databaseInterface interface {
	DeleteUserByID(id uint) error
	GetAllUsers() []model.User
	GetUserByID(id uint) (*model.User, error)
}
