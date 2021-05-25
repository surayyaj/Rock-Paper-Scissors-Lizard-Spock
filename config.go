package main

import (
	"encoding/json"
	"os"
)

type ChoiceItem struct {
	Name  string   `json:"name"`
	Beats []string `json:"beats"`
}

type Config struct {
	Choices []ChoiceItem `json:"choices"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func LoadChoices(config *Config) []*Choice {
	result := make([]*Choice, len(config.Choices))

	nameIdMap := make(map[string]int) //for faster beats populate

	for index, value := range config.Choices {
		result[index] = &Choice{ID: index + 1, Name: value.Name}
		nameIdMap[value.Name] = index + 1
	}

	// populate beats for each choice
	for index, value := range config.Choices {
		beats := []int{}
		for i := 0; i < len(value.Beats); i++ {
			name := value.Beats[i]
			id := nameIdMap[name]
			beats = append(beats, id)
		}
		result[index].Beats = beats
	}

	return result
}
