package main

import (
	"net/http"
	"printer/api/auth"
	"printer/api/jobs"
	"printer/api/middlewares"
	"printer/api/register"
	"printer/api/users"
	"printer/api/views"
	"printer/persistence/db"
	"printer/persistence/filer"
	"printer/persistence/jobq"
	"printer/persistence/session"
	"printer/worker"
)

func main() {
	// config
	server := &http.Server{Addr: "127.0.0.1:8880"}
	dsn := "host=localhost user=postgres password=password dbname=stuff port=5432 sslmode=disable"

	// persistence layer
	db := db.NewDatabase(dsn)
	jobq := jobq.NewJobQueue()
	// filer := filer.NewFiler("../files", 2<<30)
	filer := filer.NewFiler("/home/nicolas/Desktop/printer-app/files", 2<<32)
	sessionStorage := session.NewSessionStorage()
	authHandlers := auth.NewAuthHandlers(sessionStorage, db)

	// init middlewares
	requireAuth := middlewares.BuildRequireAuth(authHandlers.GetSessionIfValid, "/")
	requireNotAuth := middlewares.BuildRequireNotAuth(authHandlers.GetSessionIfValid, "/submit-file")

	// serving static files
	files := http.FileServer(http.Dir("../web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// creating views
	views := views.NewViews("../web/html")

	// authenication
	loginRedirector := middlewares.BuildRedirectOnRequest("/", "/")
	http.HandleFunc("/auth/login", requireNotAuth(loginRedirector(authHandlers.Login)))

	logoutRedirector := middlewares.BuildRedirectOnRequest("/", "/")
	http.HandleFunc("/auth/logout", requireAuth(logoutRedirector(authHandlers.Logout)))

	// set up views
	http.HandleFunc("/", requireNotAuth(views.Login))
	http.HandleFunc("/login", requireNotAuth(views.Login))
	http.HandleFunc("/signup", requireNotAuth(views.Signup))
	http.HandleFunc("/submit-file", requireAuth(views.SubmitFile))
	http.HandleFunc("/user-manager", requireAuth(views.UserManager))

	// set up jobs api routes
	jobsHandlers := jobs.NewJobsHandlers(jobq, filer)

	RedirectOnJobResult := middlewares.BuildRedirectOnRequest("/submit-file", "/user-manager")
	http.HandleFunc("/jobs/submission", requireAuth(RedirectOnJobResult(jobsHandlers.SubmitJob)))
	http.HandleFunc("/jobs/cancellation", requireAuth(RedirectOnJobResult(jobsHandlers.SubmitJob)))

	// registration
	redirectOnRegisterResult := middlewares.BuildRedirectOnRequest("/login", "/signup")
	userController := register.UserController{DB: db}
	http.HandleFunc("/register", requireNotAuth(redirectOnRegisterResult(userController.CreateNewUser)))

	usersApi := users.NewRouter(db)
	http.Handle("/users/", http.StripPrefix("/users", usersApi))

	worker := worker.NewWorker(jobq, filer, db)
	worker.Start()

	server.ListenAndServe()
}
