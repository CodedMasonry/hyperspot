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
	a.ctx = ctx
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

func (a *App) IsAuthenticated() bool {
	return a.client != nil
}

func (a *App) AuthenticateUser() string {
	config, client, err := NewSpotifyClient(a.ctx)
	if err != nil {
		log.Fatal(err)
	}

	a.config = config
	a.client = client

	return ""
}
