package main

import (
	"go-api-gin/database"
	"go-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
