package main

import (
	"log"
	"net/http"

	"github.com/marcosx3/movie-api-go/src/routes"
)

func main() {

	r := routes.Routers()

	// Start the server on port 8000
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
