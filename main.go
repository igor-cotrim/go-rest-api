package main

import (
	"fmt"

	"github.com/igor-cotrim/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
