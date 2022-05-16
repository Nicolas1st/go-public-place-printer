package account

import "printer/persistence/model"

type Database interface {
	GetUserByName(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateNewUser(username, email, password string) error
}

type accountController struct {
	db Database
}

func NewController(db Database) *accountController {
	return &accountController{db: db}
}
