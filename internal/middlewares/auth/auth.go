package middlewares

import (
	"log"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		log.Println("Auth middleware")
		next.ServeHTTP(rw, request)
	})
}
