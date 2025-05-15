package routes

import (
	"github.com/gorilla/mux"
	"github.com/marcosx3/movie-api-go/src/controller"
)

func Routers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controller.GetMovies).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.UpdateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.DeleteMovies).Methods("DELETE")
	return r
}
