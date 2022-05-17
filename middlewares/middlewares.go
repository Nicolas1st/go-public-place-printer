package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
type ApiRequestMiddleware func(func(http.ResponseWriter, *http.Request) error) http.HandlerFunc
type contextSessionKey string

var ContextSessionKey contextSessionKey = "contextSessionKey"
