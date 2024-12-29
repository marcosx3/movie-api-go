package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marcosx3/movie-api-go/pkg/director"
	"github.com/marcosx3/movie-api-go/pkg/movie"
)

var movies = []movie.Movie{
	{ID: 1, Title: "The Lord of the Rings", Director: &director.Director{FirstName: "Marcos", LastName: "Silvas"}, Rating: 9.8},
	{ID: 2, Title: "Batman The warrior shadow", Director: &director.Director{FirstName: "Ana", LastName: "leticia"}, Rating: 9.8},
	{ID: 3, Title: "The Hobbit", Director: &director.Director{FirstName: "Icaro", LastName: "Silvas"}, Rating: 8.8},
}

// Get All Movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get movie by ID
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Convert string id into int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not Founded", http.StatusNotFound)
	return
}

// Store a new Movie
func StoreMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie movie.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	movie.ID = getNextID() // necessario buscar o ultimo id disponivel
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

// Update movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedMovie movie.Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	movieID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusNotFound)
		return
	}

	// Serach the movie by ID
	for i, movie := range movies {
		if movie.ID == movieID {
			// Atualiza os dados do filme
			movies[i] = updatedMovie
			movies[i].ID = movieID // Mantém o ID original
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

// Delete a movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Convert string id into int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

// Create a new ID with the last ID +1
func getNextID() int {
	if len(movies) == 0 {
		return 1
	}
	maxID := 0
	for _, movie := range movies {
		if movie.ID > maxID {
			maxID = movie.ID
		}
	}
	return maxID + 1
}
