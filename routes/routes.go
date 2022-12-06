package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/igor-cotrim/go-rest-api/controllers"
	"github.com/igor-cotrim/go-rest-api/middlewares"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnsPersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreateNewPersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
