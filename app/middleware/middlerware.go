package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LogRequest used to log all the request
func LogRequest(router *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			defer func() {
				requestHeader := request.Header.Get("token")
				log.Printf("[%s] %s %s , %s", request.Method, request.Host, request.URL.Path, requestHeader)
			}()
			next.ServeHTTP(writer, request)
		})
	}
}

// AuthorizeRequest will check, is request authorized or not
func AuthorizeRequest(router *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			isAuth := true
			func() {
				requestHeader := request.Header.Get("Golang-X-Server")
				if requestHeader != "golangserver" {
					log.Println("authorization falied")
					isAuth = false
				}
			}()
			if !isAuth {
				http.Error(writer, "authorization falied", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(writer, request)
		})
	}
}
