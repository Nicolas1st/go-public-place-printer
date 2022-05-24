package main

import (
	"fmt"
	"net/http"
	"printer/handlers"
	"printer/handlers/api/jobs"
	"printer/handlers/api/users"
	"printer/handlers/views"
	"printer/persistence/db"
	"printer/persistence/filer"
	"printer/persistence/jobq"
	"printer/persistence/session"
	"time"
)

func main() {
	// config
	server := &http.Server{Addr: "127.0.0.1:8880"}
	dsn := "host=localhost user=postgres password=password dbname=printer port=5432 sslmode=disable"

	// set up database
	db := db.NewDatabase(dsn)
	sessioner := session.NewSessionStorage(5 * time.Minute)
	jobq := jobq.NewJobQueue()
	filer := filer.NewFiler("./files", 2<<30)

	// serve static files
	files := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// register views
	views := views.NewViews("./web/html", db, sessioner)
	http.HandleFunc(handlers.DefaultEndpoints.Root, handlers.ForNotLoggedIn(sessioner, views.Login))
	http.HandleFunc(handlers.DefaultEndpoints.LoginPage, handlers.ForNotLoggedIn(sessioner, views.Login))
	http.HandleFunc(handlers.DefaultEndpoints.SignUpPage, handlers.ForNotLoggedIn(sessioner, views.SignUp))
	http.HandleFunc(handlers.DefaultEndpoints.PrinterPage, handlers.ForCommonUsers(sessioner, views.Printer))
	http.HandleFunc(handlers.DefaultEndpoints.UserManagerPage, handlers.ForAdmin(sessioner, views.UserManager))
	http.HandleFunc(handlers.DefaultEndpoints.LogoutHandler, handlers.ForCommonUsers(sessioner, views.Logout))

	// register apis
	jobs := jobs.NewApi(jobq, filer, sessioner)
	http.HandleFunc(
		handlers.DefaultEndpoints.JobsApi,
		handlers.ForCommonUsers(sessioner, func(w http.ResponseWriter, r *http.Request) { jobs.ServeHTTP(w, r) }),
	)

	users := users.NewApi(db)
	http.HandleFunc(
		handlers.DefaultEndpoints.UsersApi,
		handlers.ForAdmin(sessioner, func(w http.ResponseWriter, r *http.Request) { users.ServeHTTP(w, r) }),
	)

	fmt.Println("Started on http://" + server.Addr)
	fmt.Println(server.ListenAndServe())
}
