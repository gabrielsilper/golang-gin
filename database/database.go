package database

import (
	"log"
	"os"

	"github.com/gabrielsilper/golang-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	dsn := "host=localhost user=root password=password dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		os.Exit(1)
	}
	log.Println("Conectado ao banco de dados")

	DB.AutoMigrate(&models.User{})
}
