package views

import (
	"net/http"
	"printer/handlers/views/pages"
	"printer/persistence/model"
	"time"
)

type database interface {
	GetUserByName(username string) (*model.User, error)
	GetUserByEmail(username string) (*model.User, error)
	CreateNewUser(username, email, password string) error
}

type sessioner interface {
	GetSessionByToken(sessionToken string) (*model.Session, bool)
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
}

type viewsController struct {
	db        database
	sessioner sessioner
}

type views struct {
	Login       http.HandlerFunc
	Logout      http.HandlerFunc
	SignUp      http.HandlerFunc
	Printer     http.HandlerFunc
	UserManager http.HandlerFunc
}

func NewViews(htmlTemplatesPath string, database database, sessioner sessioner) *views {
	c := &viewsController{
		db:        database,
		sessioner: sessioner,
	}
	pages := pages.NewPages(htmlTemplatesPath)
	return &views{
		Login:       c.buildLoginView(pages),
		Logout:      c.Logout,
		SignUp:      c.buildSignUpView(pages),
		Printer:     c.buildPrinterView(pages),
		UserManager: c.buildUserManagerView(pages),
	}
}
