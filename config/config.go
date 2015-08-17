package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port string
	Host string
}

func Load() Config {
	file, err := os.Open("../config.json")
	if err != nil {
		file, err = os.Open("config.json")
	}
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("config error:", err)
	}
	return configuration
}
