package routes

import (
	"log"
	"net/http"

	"github.com/igor-cotrim/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
