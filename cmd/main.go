package main

import (
	"fmt"
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
	dsn := "host=localhost user=postgres password=password dbname=printer port=5432 sslmode=disable"

	// persistence layer
	db := db.NewDatabase(dsn)
	jobq := jobq.NewJobQueue()

	// create admin account
	fmt.Println(db.CreateAdmin("admin", "admin@admin.admin", "admin"))

	// filer := filer.NewFiler("../files", 2<<30)
	filer := filer.NewFiler("/home/nicolas/Desktop/printer-app/files", 2<<32)
	sessionStorage := session.NewSessionStorage()
	authHandlers := auth.NewAuthHandlers(sessionStorage, db)

	// init middlewares
	requireAuth := middlewares.BuildRequireAuth(authHandlers.GetSessionIfValid, "/")
	requireNotAuth := middlewares.BuildRequireNotAuth(authHandlers.GetSessionIfValid, "/submit-file")
	onlyAdmin := middlewares.BuildOnlyAdmin("/")

	// serving static files
	files := http.FileServer(http.Dir("../web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// creating views
	views := views.NewViews("../web/html", db)

	// authenication
	loginRedirector := middlewares.BuildRedirectOnRequest("/submit-file", "/login")
	http.HandleFunc("/auth/login", requireNotAuth(loginRedirector(authHandlers.Login)))

	logoutRedirector := middlewares.BuildRedirectOnRequest("/", "/submit-file")
	http.HandleFunc("/auth/logout", requireAuth(logoutRedirector(authHandlers.Logout)))

	// set up views
	http.HandleFunc("/", requireNotAuth(views.Login))
	http.HandleFunc("/login", requireNotAuth(views.Login))
	http.HandleFunc("/signup", requireNotAuth(views.Signup))
	http.HandleFunc("/submit-file", requireAuth(views.SubmitFile))
	http.HandleFunc("/user-manager", requireAuth(onlyAdmin(views.UserManager)))
	http.HandleFunc("/profile", requireAuth(views.Profile))

	// set up jobs api routes
	jobsHandlers := jobs.NewJobsHandlers(jobq, filer)

	RedirectOnJobResult := middlewares.BuildRedirectOnRequest("/profile", "/submit-file")
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
