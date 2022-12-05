package main

import (
	"fmt"

	"github.com/igor-cotrim/go-rest-api/models"
	"github.com/igor-cotrim/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Nome 1", Story: "Historia 1"},
		{Name: "Nome 2", Story: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
