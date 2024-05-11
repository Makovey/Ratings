package models

type Movie struct {
	ID int `json:"id"`
	Name string `json:"name"`
	EnName string `json:"enName"`
	AlternativeName string `json:"alternativeName"`
	Year int `json:"year"`
}
