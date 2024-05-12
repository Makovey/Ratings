package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	m "ratings/models"

	"gopkg.in/yaml.v3"
)

const baseUrl = "https://api.kinopoisk.dev/v1.4"
const randomMovieUrl = baseUrl + "/movie/random"

var key string

func RandomMovie() m.Movie {
	client := &http.Client{}

	req, err := http.NewRequest("GET", randomMovieUrl, nil)
	req.Header.Add("X-API-KEY", getApiKey())
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("notNullFields", "name")
	
	req.URL.RawQuery = q.Encode()

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var movie m.Movie

	body, err := io.ReadAll(resp.Body)
	result := json.Unmarshal(body, &movie)
	if result != nil {
		log.Fatalln(err)
	}

	return movie
}

func getApiKey() string {
	if key != "" {
		return key
	} else {
		config := make(map[string]interface{})

		yamlFile, err := os.ReadFile("config.yaml")
		if err != nil {
			log.Fatalln(err)
		}

		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalln(err)
		}
		key = config["configuration"].(map[string]interface{})["kinopoiskApiKey"].(string)

		return key
	}
}
