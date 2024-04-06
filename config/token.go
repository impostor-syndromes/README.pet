package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token string `json:"token"`
}

func LoadToken() string {
	var config Config
	file, err := os.Open("config/token.json")
	if err != nil {
		return ""
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return ""
	}
	return config.Token
}
