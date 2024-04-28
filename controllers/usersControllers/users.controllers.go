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
		c.JSON(http.StatusBadRequest, models.NewResponseMessage(err.Error()))
		return
	}
	if err := models.ValidateUser(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseMessage(err.Error()))
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

func Update(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseMessage(err.Error()))
		return
	}
	if err := models.ValidateUser(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseMessage(err.Error()))
		return
	}
	updatedUser, err := usersService.Update(id, updatedUser)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewResponseMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := usersService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NewResponseMessage(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func SearchByName(c *gin.Context) {
	name := c.Query("name")
	users := usersService.SearchByName(name)
	c.JSON(http.StatusOK, users)
}
