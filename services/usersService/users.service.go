package usersService

import (
	"github.com/gabrielsilper/golang-gin/database"
	"github.com/gabrielsilper/golang-gin/models"
)

func Create(newUSer models.User) models.User {
	database.DB.Create(&newUSer)
	return newUSer
}

func FindAll() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}
