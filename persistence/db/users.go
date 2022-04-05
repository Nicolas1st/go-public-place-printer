package db

func (wrapper *Database) CreateNewUser(name, hashedPassword string) error {
	user := User{
		Name:         name,
		PasswordHash: hashedPassword,
	}

	result := wrapper.db.Create(&user)

	return result.Error
}

func (wrapper *Database) DeleteUserByID(id uint) error {
	result := wrapper.db.Delete(&User{}, id)

	return result.Error
}

func (wrapper *Database) GetAllUsers() []User {
	users := []User{}
	wrapper.db.Find(users)

	return users
}

func (wrapper *Database) GetUserByID(id uint) (User, error) {
	user := User{}
	result := wrapper.db.First(user, id)

	return user, result.Error
}
