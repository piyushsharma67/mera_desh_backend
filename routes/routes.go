package routes

import (
	"social_web_server/controllers"

	"github.com/gorilla/mux"
)

var P string

func InitRoutes()*mux.Router{
	r:=mux.NewRouter()

	r.HandleFunc("/",controllers.Health)
	r.HandleFunc("/home",Protected(controllers.Home))
	return r
}