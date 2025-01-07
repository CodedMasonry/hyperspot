package main

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"golang.org/x/oauth2"
)

// Functions for handling configuration file

// Calling Save is required else no config file will be written.
type Config struct {
	// Path to Config File
	path string
	// The encrypted Token.
	// To access raw token, call config.getToken()
	Token string
}

// Tries to read config.
// If no config exists, return blank config
func BuildConfig() (*Config, error) {
	configPath, err := xdg.ConfigFile("hyperspot/config.json")
	if err != nil {
		return nil, err
	}

	// If the config file exists
	if f, _ := os.Stat(configPath); f != nil {
		bytes, err := os.ReadFile(configPath)
		if err != nil {
			return nil, err
		}

		var config Config
		if err := json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}

		config.path = configPath
		return &config, nil
	}

	// If No Config Exists
	return &Config{
		path:  configPath,
		Token: "",
	}, nil
}

// Writes the config to xdg config path.
func (c *Config) Save() {
	bytes, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(c.path, bytes, 0600); err != nil {
		log.Fatal(err)
	}
}

// Saves the token; token -> json -> aes -> base64
func (c *Config) SetToken(token *oauth2.Token) error {
	// convert struct to json
	bytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	// Use path to config as encryption key so token isn't stored in plain text
	key := pathToKey(c.path)
	crypto, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// encrypt json to encryted bytes
	var encrypted []byte
	crypto.Encrypt(encrypted, bytes)

	// convert raw bytes to base64
	c.Token = base64.StdEncoding.EncodeToString(encrypted)
	return nil
}

// Retrieve the token; base64 -> aes -> json -> token
func (c *Config) GetToken() (*oauth2.Token, error) {
	// Check if token has been set
	if c.Token == "" {
		return nil, fmt.Errorf("no token set; please set a token before trying to read one")
	}

	// Convert base64 token to raw bytes
	bytes := make([]byte, base64.StdEncoding.DecodedLen(len(c.Token)))
	_, err := base64.StdEncoding.Decode(bytes, []byte(c.Token))
	if err != nil {
		return nil, err
	}

	// Use path to config as encryption key so token isn't stored in plain text
	key := pathToKey(c.path)
	crypto, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// decrypt bytes to json
	var decrypted []byte
	crypto.Decrypt(decrypted, []byte(c.Token))

	// convert json to token
	var token *oauth2.Token
	err = json.Unmarshal(decrypted, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Don't hardcode a key, just derive one from config path
func pathToKey(path string) []byte {
	bytes := make([]byte, 32)

	// Take every character
	for idx, char := range path {
		if idx < 32 {
			// set byte to char
			bytes[idx] = byte(char)
		} else {
			// xor previous char with current char
			bytes[idx] = bytes[idx] ^ byte(char)
		}
	}

	return bytes
}
