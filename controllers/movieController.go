package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	m "ratings/models"

	"gopkg.in/yaml.v3"
)

// TODO: handle error with logger
func RandomMovie() *m.Movie {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.kinopoisk.dev/v1.4/movie/random", nil)
	req.Header.Add("X-API-KEY", getApiKey())
	req.Header.Add("Accept", "application/json")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	var movie m.Movie

	body, err := io.ReadAll(resp.Body)
	result := json.Unmarshal(body, &movie)
	if result != nil {
		fmt.Println(err)
		return nil
	}

	return &movie
}

// TODO: readfile one time
// TODO: change to struct
func getApiKey() string {
	config := make(map[string]interface{})

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return config["configuration"].(map[string]interface{})["kinopoiskApiKey"].(string)
}
