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
	StoreSession(session *model.Session) (string, time.Time)
	RemoveSession(sessionToken string)
}

type viewsController struct {
	db        database
	sessioner sessioner
}

type views struct {
	Login       http.HandlerFunc
	SignUp      http.HandlerFunc
	SubmitFile  http.HandlerFunc
	UserManager http.HandlerFunc
}

func NewViews(htmlTemplatesPath string, database database, sessioner sessioner) *views {
	c := &viewsController{
		db:        database,
		sessioner: sessioner,
	}
	pages := pages.NewPages(htmlTemplatesPath)
	return &views{
		Login:       c.buildLoginView(pages.Login),
		SignUp:      c.buildSignUpView(pages.Signup),
		SubmitFile:  buildView(pages.SubmitFile),
		UserManager: buildView(pages.UserManager),
	}
}
