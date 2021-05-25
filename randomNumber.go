package main

import (
	"encoding/json"
	"net/http"
)

type RandomNumberResponse struct {
	RandomNumber int `json:"random_number"`
}

func getRandomNumber() int {
	//return 2
	r, _ := http.Get(config.RandomNumberUrl)

	var response RandomNumberResponse
	_ = json.NewDecoder(r.Body).Decode(&response)

	return response.RandomNumber
}
