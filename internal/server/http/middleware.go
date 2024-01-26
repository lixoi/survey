package internalhttp

import (
	"net/http"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc { //nolint:unused
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		next(w, r)
	}
}
