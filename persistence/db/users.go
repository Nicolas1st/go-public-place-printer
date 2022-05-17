package db

import (
	"errors"
	"fmt"
	"printer/persistence/model"

	"golang.org/x/crypto/bcrypt"
)

func (wrapper *Database) CreateAdmin(name, email, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		fmt.Println("Could not create the admin account")
		return err
	}

	user := model.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
		Role:         model.Admin,
	}

	result := wrapper.db.Create(&user)

	return result.Error
}

// CreateNewUser - creates new user, if it's not possible an error value is returned
// password hashing is performed by this function
func (wrapper *Database) CreateNewUser(name, email, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return errors.New("could not hash the password")
	}

	user := model.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
		Role:         model.NonAdmin,
	}

	return wrapper.db.Create(&user).Error
}

// DeleteUserByID - deletes user by id
func (wrapper *Database) DeleteUserByID(id uint) error {
	result := wrapper.db.Delete(&model.User{}, id)

	return result.Error
}

// GetUserByID - retrieve one user by id
func (wrapper *Database) GetUserByID(id uint) (*model.User, error) {
	user := model.User{}
	result := wrapper.db.First(&user, id)

	return &user, result.Error
}

// GetUserByName - retrieve one user by name
func (wrapper *Database) GetUserByName(username string) (*model.User, error) {
	user := model.User{}
	result := wrapper.db.Where("Name = ?", username).First(&user)

	return &user, result.Error
}

// GetUserByEmail - retrieve one user by email
func (wrapper *Database) GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	result := wrapper.db.Where("Email = ?", email).First(&user)

	return &user, result.Error
}

// GetAllUsers - retrieves all users from database
func (wrapper *Database) GetAllUsers() []model.User {
	users := []model.User{}
	wrapper.db.Find(&users)

	return users
}
