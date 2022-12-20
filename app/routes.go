package app

import (
	"github.com/alwinre-factory/GoToko/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) InitializeRouters() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
