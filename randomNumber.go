package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(response.RandomNumber)

	return response.RandomNumber
}
