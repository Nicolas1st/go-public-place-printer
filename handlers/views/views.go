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
	GetAllPrints() []model.Print
	GetAllPrintsByUID(UID uint) []model.Print
	GetPrintsForDayNDaysAgo(daysAgo int) ([]model.Print, error)
	GetAllPrintsByUsername(username string) []model.Print
}

type sessioner interface {
	GetSessionByToken(sessionToken string) (*model.Session, bool)
	StoreSession(user *model.User) (string, time.Time)
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
	Stats       http.HandlerFunc
	Prints      http.HandlerFunc
	UserPrints  http.HandlerFunc
	FileRemoved http.HandlerFunc
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
		Stats:       c.buildStatsView(pages),
		Prints:      c.buildPrintsView(pages),
		UserPrints:  c.buildUserPrintsView(pages),
		FileRemoved: c.buildFileRemovedView(pages),
	}
}
