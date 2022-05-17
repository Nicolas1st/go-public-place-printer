package main

import (
	"net/http"
	"printer/handlers"
	"printer/handlers/views"
	"printer/persistence/db"
	"printer/persistence/session"
)

func main() {
	// config
	server := &http.Server{Addr: "127.0.0.1:8880"}
	dsn := "host=localhost user=postgres password=password dbname=printer port=5432 sslmode=disable"

	// set up database
	db := db.NewDatabase(dsn)
	sessioner := session.NewSessionStorage()

	// serve static files
	files := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// register views
	views := views.NewViews("./web/html", db, sessioner)
	http.Handle(handlers.DefaultEndpoints.Root, handlers.ForNotLoggedIn(sessioner, views.Login))
	http.Handle(handlers.DefaultEndpoints.LoginPage, handlers.ForNotLoggedIn(sessioner, views.Login))
	http.Handle(handlers.DefaultEndpoints.SignUpPage, handlers.ForNotLoggedIn(sessioner, views.SignUp))
	http.Handle(handlers.DefaultEndpoints.PrinterPage, handlers.ForCommonUsers(sessioner, views.Printer))
	http.Handle(handlers.DefaultEndpoints.UserManagerPage, handlers.ForAdmin(sessioner, views.UserManager))

	server.ListenAndServe()
}
