package usersControllers

import (
	"net/http"

	"github.com/gabrielsilper/golang-gin/services/usersService"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	users := usersService.FindAll()
	c.JSON(http.StatusOK, users)
}
