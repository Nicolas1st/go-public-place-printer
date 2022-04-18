package users

type createUserRequestBody struct {
	Name         string `json:"Name"`
	PasswordHash string `json:"PasswordHash"`
}

type deleteUserByIDRequestBody struct {
	UID uint `json:"UID"`
}

type getUserByIDRequestBody struct {
	UID uint `json:"UID"`
}
