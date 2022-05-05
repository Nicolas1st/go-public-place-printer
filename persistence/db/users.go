package db

import "printer/persistence/model"

// CreateNewUser - creates new user, if it's not possible an error value is returned
func (wrapper *Database) CreateNewUser(name, email, hashedPassword string) error {
	user := model.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	result := wrapper.db.Create(&user)

	return result.Error
}

func (wrapper *Database) DeleteUserByID(id uint) error {
	result := wrapper.db.Delete(&model.User{}, id)

	return result.Error
}

func (wrapper *Database) GetAllUsers() []model.User {
	users := []model.User{}
	wrapper.db.Find(users)

	return users
}

func (wrapper *Database) GetUserByID(id uint) (model.User, error) {
	user := model.User{}
	result := wrapper.db.First(user, id)

	return user, result.Error
}

func (wrapper *Database) GetUserByName(username string) (*model.User, error) {
	user := model.User{}
	result := wrapper.db.Where("Name = ?", username).First(&user)

	return &user, result.Error
}
