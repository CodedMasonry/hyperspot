package main

// Functions for handling communication between the Frontend & the Backend

import (
	"context"
	"log"

	"github.com/zmb3/spotify/v2"
)

// App struct
type App struct {
	ctx    context.Context
	config *Config
	client *spotify.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	config, client, err := NewSpotifyClient(a.ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Save state
	a.config = config
	a.ctx = ctx
	// If user has logged in, client will be valid
	// If not, client is null & login is expected
	a.client = client
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Whether the user is logged in already
func (a *App) IsAuthenticated() bool {
	return a.client != nil
}

// Login user if they aren't
func (a *App) AuthenticateUser() bool {
	// Set client since user logged in
	a.client = LoginSpotify(a.ctx)

	// pull token from client
	token, err := a.client.Token()
	if err != nil {
		log.Fatalf("failed to get token: %v\n", err)
	}
	a.config.SetToken(token)

	// Manually save config
	a.config.Save()
	// Return a bool simply to notify frontend that authentication is complete
	return true
}
