package app

import "github.com/Yaantirahmaw/GoStore/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

}