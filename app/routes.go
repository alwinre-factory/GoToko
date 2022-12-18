package app

import "github.com/alwinre-factory/GoToko/app/controllers"

func (server *Server) InitializeRouters() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
