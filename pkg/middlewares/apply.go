package middlewares

import "net/http"

func ApllyMiddlewares(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	// return unchange if there are not middlewares
	if len(m) < 1 {
		return h
	}

	wrapped := h
	// apply middlwares in reverse order, to keep correst nesting
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}
