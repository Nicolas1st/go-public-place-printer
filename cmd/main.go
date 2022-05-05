package main

import (
	"net/http"
	"printer/api/auth"
	"printer/api/jobs"
	"printer/api/middlewares"
	"printer/api/register"
	"printer/api/views"
	"printer/persistence/db"
	"printer/persistence/filer"
	"printer/persistence/jobq"
	"printer/persistence/session"
	"printer/worker"
)

func main() {
	// config
	server := &http.Server{
		Addr: "127.0.0.1:8880",
	}
	dsn := "host=localhost user=postgres password=password dbname=stuff port=5432 sslmode=disable"

	// persistence layer
	db := db.NewDatabase(dsn)
	jobq := jobq.NewJobQueue()
	filer := filer.NewFiler("../files", 2<<30)
	sessionStorage := session.NewSessionStorage()

	// serving static files
	files := http.FileServer(http.Dir("../web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// creating views
	views := views.NewViews("../web/html")

	// build middlewares for http handlers
	onlyAuth := middlewares.BuildOnlyAuthenticatedMiddleware(sessionStorage, views.Login)
	onlyNotAuth := middlewares.BuildOnlyAnonymousMiddleware(views.SubmitFile)

	// setting up views
	// not protected routes
	http.HandleFunc("/", onlyNotAuth(views.Login))
	http.HandleFunc("/login", onlyNotAuth(views.Login))
	http.HandleFunc("/signup", onlyNotAuth(views.Signup))

	// protected routes
	http.HandleFunc("/submit-file", onlyAuth(views.SubmitFile))

	// jobs api to submit files to be printed
	jobsApi := jobs.NewRouter(jobq, filer)
	http.Handle("/jobs", onlyAuth(jobsApi))

	// registration
	userController := register.UserController{
		DB: db,
	}
	http.HandleFunc("/register", userController.CreateNewUser)

	// authenication
	auth := auth.NewAuthRouter(
		sessionStorage,
		db,
		views.SubmitFile,
		views.Login,
	)
	http.Handle("/auth/", http.StripPrefix("/auth", auth))

	worker := worker.NewWorker(jobq, filer)
	worker.Start()

	server.ListenAndServe()
}
