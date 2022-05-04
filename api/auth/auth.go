package auth

import "net/http"

type AuthResource struct {
	sessionStorage SessionStorageInterface
	database       DatabaseInterface
}

func newAuthResource(sessionStorage SessionStorageInterface, database DatabaseInterface) *AuthResource {
	return &AuthResource{
		sessionStorage: sessionStorage,
		database:       database,
	}
}

func NewAuthRouter(
	sessionStorage SessionStorageInterface,
	database DatabaseInterface,
	redirectToOnLogin http.HandlerFunc,
	redirectToOnLogout http.HandlerFunc,
) *http.ServeMux {
	authResource := newAuthResource(sessionStorage, database)

	router := http.NewServeMux()
	router.HandleFunc("/login", authResource.buildAuthenticate(redirectToOnLogin))
	router.HandleFunc("/logout", authResource.buildLogout(redirectToOnLogout))

	return router
}
