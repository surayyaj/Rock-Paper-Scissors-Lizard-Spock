package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var config Config

var choices []*Choice

func main() {
	config, _ = LoadConfiguration("config.json")
	choices = LoadChoices(&config)

	r := mux.NewRouter()
	r.HandleFunc("/choices", getChoices).Methods("GET")
	r.HandleFunc("/choice", getChoice).Methods("GET")
	r.HandleFunc("/play", play).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(config.Host, r))
}

func getChoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(choices)
}

func getChoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	randomNumber := getRandomNumber()
	json.NewEncoder(w).Encode(choices[randomNumber%len(choices)])
}

func play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	var request PlayRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	computerChoiceId := getRandomNumber()%len(choices) + 1
	results := PlayGame(request.Player, computerChoiceId, choices)
	response := PlayResponse{
		Results:  results,
		Player:   request.Player,
		Computer: computerChoiceId,
	}
	json.NewEncoder(w).Encode(response)
}
