package movie

import "github.com/marcosx3/movie-api-go/pkg/director"

type Movie struct {
	ID       int                `json:"id"`
	Title    string             `json:"title"`
	Director *director.Director `json:"director"`
	Rating   float64            `json:"rating"`
}
