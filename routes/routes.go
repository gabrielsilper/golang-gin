package routes

import (
	"net/http"

	"github.com/gabrielsilper/golang-gin/controllers/usersControllers"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartServer() {
	router.GET("/live", LiveApi)
	handleUsersRequests()
	router.Run()
}

func handleUsersRequests() {
	usersRouter := router.Group("/api/users")
	usersRouter.GET("/", usersControllers.FindAll)
	usersRouter.GET("/:id", usersControllers.FindById)
	usersRouter.GET("/search", usersControllers.SearchByName)
	usersRouter.POST("/", usersControllers.Create)
	usersRouter.PUT(":id", usersControllers.Update)
	usersRouter.DELETE(":id", usersControllers.Delete)
}

func LiveApi(c *gin.Context) {
	c.String(http.StatusOK, "Server is live...")
}
