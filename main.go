package main

import (
	"fmt"

	"github.com/igor-cotrim/go-rest-api/database"
	"github.com/igor-cotrim/go-rest-api/models"
	"github.com/igor-cotrim/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Nome 1", Story: "Historia 1"},
		{Id: 2, Name: "Nome 2", Story: "Historia 2"},
	}
	database.ConnectToDatabase()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
