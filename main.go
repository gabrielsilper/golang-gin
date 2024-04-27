package main

import (
	"github.com/gabrielsilper/golang-gin/database"
	"github.com/gabrielsilper/golang-gin/routes"
)

func main() {
	database.ConnectToDB()
	routes.StartServer()
}
