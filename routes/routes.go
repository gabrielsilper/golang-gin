package routes

import (
	"github.com/gabrielsilper/golang-gin/controllers/usersControllers"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartServer() {
	handleUsersRequests()
	router.Run()
}

func handleUsersRequests() {
	apiRouter := router.Group("/api/users")
	apiRouter.GET("/", usersControllers.FindAll)
	apiRouter.POST("/", usersControllers.Create)
}
