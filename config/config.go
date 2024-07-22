package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DatabaseFile string `json:"database_file"`
	KeyFile      string `json:"key_file"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
