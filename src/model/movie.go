package model

type Movie struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Rating   float32   `json:"rating"`
	Director *Director `json:"director"`
}
