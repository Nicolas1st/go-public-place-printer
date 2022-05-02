package register

type DatabaseInterface interface {
	CreateNewUser(username, email, passwordHash string) error
}
