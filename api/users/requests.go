package users

type createUserRequestBody struct {
	name         string `json:"name"`
	passwordHash string `json:"passwordHash"`
}

type deleteUserByIDRequestBody struct {
	UID uint `json:"UID"`
}

type getUserByIDRequestBody struct {
	UID uint `json:"UID"`
}
