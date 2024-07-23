package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	PlayerName string `json:"player_name"`
}

const defaultConfigPath = "config.json"

func Load() (*Config, error) {
	// Check if config file exists, create default if it doesn't
	if _, err := os.Stat(defaultConfigPath); os.IsNotExist(err) {
		return createDefaultConfig()
	}

	// Read and parse the config file
	file, err := os.Open(defaultConfigPath)
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

func createDefaultConfig() (*Config, error) {
	config := &Config{
		PlayerName: "Player",
	}

	file, err := os.Create(defaultConfigPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
