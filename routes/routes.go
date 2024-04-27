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
	usersRouter := router.Group("/api/users")
	usersRouter.GET("/", usersControllers.FindAll)
	usersRouter.GET("/:id", usersControllers.FindById)
	usersRouter.POST("/", usersControllers.Create)
	usersRouter.PUT(":id", usersControllers.Update)
	usersRouter.DELETE(":id", usersControllers.Delete)
}
