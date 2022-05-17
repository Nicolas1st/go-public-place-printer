package config

type Endpoints struct {
	// Public
	LoginPage            string
	SignUpPage           string
	CreateAccountHandler string

	// Private
	LogoutHandler         string
	ProfilePage           string
	PrinterPage           string
	SubmitJobHandler      string
	CancelJobHandler      string
	UpdateEmailHandler    string
	UpdatePasswordHandler string

	// Admin
	UserManagerPage string
}

var DefaultEndpoints = Endpoints{
	LoginPage:            "/login",
	SignUpPage:           "/signup",
	CreateAccountHandler: "/account/creation",

	LogoutHandler:         "/logout",
	ProfilePage:           "/profile",
	PrinterPage:           "/printer",
	SubmitJobHandler:      "/printer/jobs",
	CancelJobHandler:      "/printer/jobs",
	UpdateEmailHandler:    "/account/email",
	UpdatePasswordHandler: "/account/password",

	UserManagerPage: "/admin/users",
}
