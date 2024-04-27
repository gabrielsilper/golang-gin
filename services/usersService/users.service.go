package usersService

import (
	"errors"
	"fmt"

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

func FindById(id string) (models.User, error) {
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		errorMessage := fmt.Sprintf("Usuário de ID #%s não encontrado!", id)
		return user, errors.New(errorMessage)
	}
	return user, nil
}
