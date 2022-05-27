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
	PrinterPage           string
	SubmitJobHandler      string
	CancelJobHandler      string
	UpdateEmailHandler    string
	UpdatePasswordHandler string
	JobsApi               string
	UsersApi              string
	AccountsApi           string
	StatsApi              string

	// Admin
	UserManagerPage string
	StatsPage       string
	PrintsPage      string
	UserPrintsPage  string
	FileRemovedPage string
}

var DefaultEndpoints = Endpoints{
	Root:              "/",
	LoginPage:         "/login",
	LoginFormHandler:  "/login",
	SignUpPage:        "/signup",
	SignUpFormHandler: "/signup",

	LogoutHandler: "/logout",
	PrinterPage:   "/printer",

	UpdateEmailHandler:    "/account/email",
	UpdatePasswordHandler: "/account/password",

	JobsApi:          "/printer/jobs/",
	SubmitJobHandler: "/printer/jobs/",
	CancelJobHandler: "/printer/jobs/",

	UsersApi:    "/users/",
	AccountsApi: "/accounts/",
	StatsApi:    "/stats/",

	UserManagerPage: "/admin/users",
	StatsPage:       "/stats",
	PrintsPage:      "/prints",
	UserPrintsPage:  "/user-prints",
	FileRemovedPage: "/file/status/removed",
}
