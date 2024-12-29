package main

import (
	"log"
	"net/http"

	"github.com/marcosx3/movie-api-go/internal/routes"
)

func main() {

	r := routes.SetupRoutes() // Configurando as rotas

	log.Println("Servidor rodando na porta 5000")

	if err := http.ListenAndServe(":5000", r); err != nil {

		log.Fatal(err)

	}
}
