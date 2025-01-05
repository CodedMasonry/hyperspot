package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/adrg/xdg"
)

// Functions for handling configuration file

// Calling Save is required else no config file will be written.
type Config struct {
	// Path to Config File
	path string
	// Refresh token for Spotify Client
	RefreshToken string
}

// Tries to read config.
// If no config exists, return blank config
func GetConfig() *Config {
	configPath, err := xdg.ConfigFile("hyperspot/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// If the config file exists
	if f, _ := os.Stat(configPath); f != nil {
		bytes, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatal(err)
		}

		var config Config
		if err := json.Unmarshal(bytes, &config); err != nil {
			log.Fatal(err)
		}

		config.path = configPath
		return &config
	}
	// If No Config Exists
	return &Config{
		path: configPath,
	}
}

func (c *Config) Save() {
	bytes, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(c.path, bytes, 0600); err != nil {
		log.Fatal(err)
	}
}
