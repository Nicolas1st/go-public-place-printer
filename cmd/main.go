package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"printer/handlers"
	"printer/handlers/api/jobs"
	"printer/handlers/api/stats"
	"printer/handlers/api/users"
	"printer/handlers/views"
	"printer/persistence/db"
	"printer/persistence/filer"
	"printer/persistence/jobq"
	"printer/persistence/session"
	"printer/pkg/dirsize"
	"time"
)

func getEnvVar(envVarName string) string {
	envVar := os.Getenv(envVarName)
	if envVar == "" {
		panic("must provide " + envVarName)
	}

	return envVar
}

func main() {
	// read in the env vars
	APP_DB_USER := getEnvVar("APP_DB_USER")
	APP_DB_HOST := getEnvVar("APP_DB_HOST")
	APP_DB_PASSWORD := getEnvVar("APP_DB_PASSWORD")
	APP_DB_DBNAME := getEnvVar("APP_DB_DBNAME")
	APP_DB_PORT := getEnvVar("APP_DB_PORT")
	APP_PATH_TO_FILES := getEnvVar("APP_PATH_TO_FILES")
	APP_PORT := getEnvVar("APP_PORT")
	APP_ADMIN_LOGIN := getEnvVar("admin")
	APP_ADMIN_EMAIL := getEnvVar("admin@admin.admin")
	APP_ADMIN_PASSWORD := getEnvVar("adminpassword")

	// config
	server := &http.Server{Addr: ":" + APP_PORT}

	// set up storage
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", APP_DB_HOST, APP_DB_USER, APP_DB_PASSWORD, APP_DB_DBNAME, APP_DB_PORT)
	db := db.NewDatabase(dsn)
	filer := filer.NewFiler(APP_PATH_TO_FILES, 2<<30)
	sessioner := session.NewSessionStorage(5 * time.Minute)
	jobq := jobq.NewJobQueue()

	// создать аккаунт админа, если еще не создан
	db.CreateAdmin(APP_ADMIN_LOGIN, APP_ADMIN_EMAIL, APP_ADMIN_PASSWORD)

	// serve static files
	files := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// register views
	views := views.NewViews("./web/html", db, sessioner)

	// public
	http.HandleFunc(handlers.DefaultEndpoints.Root, handlers.ForNotLoggedIn(sessioner, views.Login))
	http.HandleFunc(handlers.DefaultEndpoints.LoginPage, handlers.ForNotLoggedIn(sessioner, views.Login))
	http.HandleFunc(handlers.DefaultEndpoints.SignUpPage, handlers.ForNotLoggedIn(sessioner, views.SignUp))
	// private
	http.HandleFunc(handlers.DefaultEndpoints.PrinterPage, handlers.ForCommonUsers(sessioner, views.Printer))
	http.HandleFunc(handlers.DefaultEndpoints.LogoutHandler, handlers.ForCommonUsers(sessioner, views.Logout))
	// admin
	http.HandleFunc(handlers.DefaultEndpoints.UserManagerPage, handlers.ForAdmin(sessioner, views.UserManager))
	http.HandleFunc(handlers.DefaultEndpoints.StatsPage, handlers.ForAdmin(sessioner, views.Stats))
	http.HandleFunc(handlers.DefaultEndpoints.PrintsPage, handlers.ForAdmin(sessioner, views.Prints))
	http.HandleFunc(handlers.DefaultEndpoints.UserPrintsPage, handlers.ForAdmin(sessioner, views.UserPrints))
	http.HandleFunc(handlers.DefaultEndpoints.FileRemovedPage, handlers.ForAdmin(sessioner, views.FileRemoved))

	// register apis
	jobs := jobs.NewApi(jobq, filer, sessioner, db)
	http.HandleFunc(
		handlers.DefaultEndpoints.JobsApi,
		handlers.ForCommonUsers(sessioner, func(w http.ResponseWriter, r *http.Request) { jobs.ServeHTTP(w, r) }),
	)

	users := users.NewApi(db)
	http.Handle(
		handlers.DefaultEndpoints.UsersApi,
		handlers.ForAdmin(sessioner, func(w http.ResponseWriter, r *http.Request) { users.ServeHTTP(w, r) }),
	)

	stats := stats.NewApi(db)
	http.HandleFunc(
		handlers.DefaultEndpoints.StatsApi,
		handlers.ForAdmin(sessioner, func(w http.ResponseWriter, r *http.Request) { stats.ServeHTTP(w, r) }),
	)

	// serve printed files
	printedFiles := http.FileServer(http.Dir(APP_PATH_TO_FILES))
	http.HandleFunc(
		APP_PATH_TO_FILES+"/",
		handlers.ForAdmin(
			sessioner,
			func(w http.ResponseWriter, r *http.Request) {
				if _, err := os.Stat(r.URL.Path); err != nil {
					http.Redirect(w, r, handlers.DefaultEndpoints.FileRemovedPage, http.StatusSeeOther)
					return
				}

				http.StripPrefix(APP_PATH_TO_FILES, printedFiles).ServeHTTP(w, r)
			},
		),
	)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			gbUsed, err := dirsize.DirSize(APP_PATH_TO_FILES)
			if err != nil {
				continue
			}
			fmt.Println(gbUsed)
			if gbUsed > 1 {
				log.Println(os.RemoveAll(APP_PATH_TO_FILES))
				log.Println(os.Mkdir(APP_PATH_TO_FILES, os.ModeAppend))
			}
		}
	}()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			jobq.RemoveCompletedAndCanceledJobs()
		}
	}()

	fmt.Println("Started on http://" + server.Addr)
	fmt.Println(server.ListenAndServe())
}
