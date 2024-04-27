package usersService

import "github.com/gabrielsilper/golang-gin/models"

func FindAll() []models.User {
	users := []models.User{
		{ID: 1, Name: "gabrielsilper", Email: "gabrielsilper@gmail.com"},
		{ID: 2, Name: "fulano", Email: "fulano@gmail.com"},
	}
	return users
}
