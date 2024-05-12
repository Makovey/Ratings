package models

type Movie struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	AlternativeName string `json:"alternativeName"`
	Year            int    `json:"year"`
	Description     string `json:"description"`
	Rating          Rating `json:"rating"`
	Poster          Poster `json:"poster"`
}

type Rating struct {
	Kp   float32 `json:"kp"`
	Imdb float32 `json:"imdb"`
}

type Poster struct {
	Url string `json:"url"`
}
