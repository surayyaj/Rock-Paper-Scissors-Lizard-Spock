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
	r, _ := http.Get("https://codechallenge.boohma.com/random")

	var response RandomNumberResponse
	_ = json.NewDecoder(r.Body).Decode(&response)

	fmt.Println(response.RandomNumber)

	return response.RandomNumber
}
