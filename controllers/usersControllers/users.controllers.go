package usersControllers

import (
	"net/http"

	"github.com/gabrielsilper/golang-gin/models"
	"github.com/gabrielsilper/golang-gin/services/usersService"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser = usersService.Create(newUser)
	c.JSON(http.StatusCreated, newUser)
}

func FindAll(c *gin.Context) {
	users := usersService.FindAll()
	c.JSON(http.StatusOK, users)
}

func FindById(c *gin.Context) {
	id := c.Param("id")
	user, err := usersService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewResponseMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, user)
}
