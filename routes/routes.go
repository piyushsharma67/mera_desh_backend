package routes

import (
	"fmt"
	"net/http"
	controllers "social_web_server/handler"
	"time"

	"github.com/gorilla/mux"
)

var P string

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func InitRoutes()*mux.Router{
	r:=mux.NewRouter()
	r.Use(LoggingMiddleware)

	r.HandleFunc("/",controllers.Health)
	r.HandleFunc("/signup",controllers.SignupUser).Methods("POST")
	return r
}