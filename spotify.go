package main

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

// Package for functions related to Spotify API

func NewSpotifyClient(ctx context.Context) (*Config, *spotify.Client, error) {
	config, err := BuildConfig()
	if err != nil {
		return nil, nil, err
	}

	var client *spotify.Client

	if config.Token == "" {
		client = LoginSpotify(ctx)

		token, err := client.Token()
		if err != nil {
			return nil, nil, err
		}

		config.SetToken(token)
	} else {
		token, err := config.GetToken()
		if err != nil {
			return nil, nil, err
		}
		client = spotify.New(auth.Client(ctx, token))
	}

	return config, client, nil
}
