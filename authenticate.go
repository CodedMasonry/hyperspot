package main

// Functions for handling Login & Authentication of Users
// Initial Login & Refresh token handling is done here

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const redirectURI = "http://localhost:6873/callback"

// Client ID for Spotify Project
const SpotifyClientId = "4c90884311b14297ac933a092eadcbb8"

var (
	auth              = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	authChannel       = make(chan *spotify.Client)
	authState         = "hyperspot"
	authCodeVerifier  = oauth2.GenerateVerifier()
	authCodeChallenge = oauth2.S256ChallengeFromVerifier(authCodeVerifier)
)

func LoginSpotify(ctx context.Context) *spotify.Client {
	http.HandleFunc("/callback", completeAuth)
	go http.ListenAndServe(":6873", nil)

	url := auth.AuthURL(authState,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", authCodeChallenge),
		oauth2.SetAuthURLParam("client_id", SpotifyClientId))

	runtime.BrowserOpenURL(ctx, url)

	client := <-authChannel
	return client
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), authState, r, oauth2.SetAuthURLParam("code_verifier", authCodeVerifier), oauth2.SetAuthURLParam("client_id", SpotifyClientId))
	if err != nil {
		http.Error(w, "Couldn't Get Token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != authState {
		http.NotFound(w, r)
		log.Fatalf("State Mismatch: %s != %s\n", st, authState)
	}

	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed! You Can Close This Tab")
	authChannel <- client
}
