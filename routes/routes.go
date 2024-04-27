package routes

import (
	"github.com/gabrielsilper/golang-gin/controllers/usersControllers"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartServer() {
	handleRequests()
	router.Run()
}

func handleRequests() {
	apiRouter := router.Group("/api")
	apiRouter.GET("/alunos", usersControllers.FindAll)
}
