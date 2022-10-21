package main

import (
	"project-pertama/sesi-10/server"
	"project-pertama/sesi-10/server/controller"
)

func main() {

	userController := controller.NewUserController()
	router := server.NewGinRouter(":4444", userController)

	router.Start()
}
