package handlers

type Endpoints struct {
	Root string

	// Public
	LoginPage         string
	LoginFormHandler  string
	SignUpPage        string
	SignUpFormHandler string

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
	Root:              "/",
	LoginPage:         "/login",
	LoginFormHandler:  "/login",
	SignUpPage:        "/signup",
	SignUpFormHandler: "/signup",

	LogoutHandler:         "/logout",
	ProfilePage:           "/profile",
	PrinterPage:           "/printer",
	SubmitJobHandler:      "/printer/jobs",
	CancelJobHandler:      "/printer/jobs",
	UpdateEmailHandler:    "/account/email",
	UpdatePasswordHandler: "/account/password",

	UserManagerPage: "/admin/users",
}
