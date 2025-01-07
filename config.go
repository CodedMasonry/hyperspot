package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"golang.org/x/crypto/chacha20poly1305"
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

// Saves the token; token -> json -> encrypt -> hex
func (c *Config) SetToken(token *oauth2.Token) error {
	// convert struct to json
	bytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	// Use path to config as encryption key so token isn't stored in plain text
	key := pathToKey(c.path)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}

	// Encryption nonce
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(bytes)+aead.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	// encrypt json to encryted bytes
	encrypted := aead.Seal(nonce, nonce, bytes, nil)

	// convert raw bytes to hex
	result := hex.EncodeToString(encrypted)
	c.Token = result
	return nil
}

// Retrieve the token; hex -> decrypt -> json -> token
func (c *Config) GetToken() (*oauth2.Token, error) {
	// Check if token has been set
	if c.Token == "" {
		return nil, fmt.Errorf("no token set; please set a token before trying to read one")
	}

	// Convert hex token to raw bytes
	bytes, err := hex.DecodeString(c.Token)
	if err != nil {
		return nil, err
	}

	// Use path to config as encryption key so token isn't stored in plain text
	key := pathToKey(c.path)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}

	// Decrypt to raw bytes
	nonce, ciphertext := bytes[:aead.NonceSize()], bytes[aead.NonceSize():]
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	// convert json to token
	var token oauth2.Token
	err = json.Unmarshal(plaintext, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
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
			pos := idx % 32
			// xor previous char with current char
			bytes[pos] = bytes[pos] ^ byte(char)
		}
	}

	return bytes
}
