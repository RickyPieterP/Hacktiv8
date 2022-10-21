package server

import (
	"project-pertama/sesi-10/server/controller"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	port string
	user *controller.UserController
}

func NewGinRouter(port string, user *controller.UserController) *GinRouter {
	return &GinRouter{
		port: port,
		user: user,
	}
}

func (g *GinRouter) Start() {
	router := gin.Default()
	router.POST("/users/register", g.user.CreateUser)
	router.POST("/users/login", g.user.Login)
	router.Run(g.port)
}
