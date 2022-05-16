package account

import "printer/persistence/model"

type database interface {
	GetUserByName(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateNewUser(username, email, password string) error
}

type accountController struct {
	db database
}

func NewController(db database) *accountController {
	return &accountController{db: db}
}
