package routes

import (
	"github.com/gorilla/mux"
	"github.com/marcosx3/movie-api-go/internal/controllers"
)

// Configura as rotas da aplicação
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controllers.StoreMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	return r
}
