package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marcosx3/movie-api-go/src/model"
)

var movies []model.Movie

// movies = append(movies, model.Movie{ID: 1, Name: "The Shawshank Redemption", Rating: 9.3, Director: &model.Director{ID: 1, Name: "Frank Darabont"}})
// movies = append(movies, model.Movie{ID: 2, Name: "The Godfather", Rating: 9.2, Director: &model.Director{ID: 2, Name: "Francis Ford Coppola"}})

// GetMovies returns all movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

// GetMovie returns a movie by ID
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreateMovies creates a new movie
func CreateMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = len(movies) + 1
	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

// UpdateMovies updates an existing movie
func UpdateMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			var movie model.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = id
			movies = append(movies, movie)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// DeleteMovies deletes a movie by ID
func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
