package usersControllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}
