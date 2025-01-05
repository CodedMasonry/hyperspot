package main

// Functions for handling Login & Authentication of Users
// Initial Login & Refresh token handling is done here

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/exp/rand"
	"golang.org/x/oauth2"
)

const redirectURI = "http://localhost:6873"
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var (
	auth          = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	ch            = make(chan *spotify.Client)
	state         = "hyperspot"
	codeVerifier  = generateRandomString(64)
	codeChallenge = generateCodeChallenge(codeVerifier)
)

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func generateCodeChallenge(code string) string {
	hash := sha256.Sum256([]byte(code))

	dst := make([]byte, base64.URLEncoding.Strict().EncodedLen(len(hash)))
	base64.URLEncoding.Encode(dst, hash[:])

	return string(dst)
}

func HandleLogin(ctx context.Context) *spotify.Client {
	http.HandleFunc("/callback", completeAuth)
	go http.ListenAndServe(":6873", nil)

	url := auth.AuthURL(state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("client_id", os.Getenv("CLIENT_ID")))

	runtime.BrowserOpenURL(ctx, url)

	client := <-ch
	return client
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		http.Error(w, "Couldn't Get Token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State Mismatch: %s != %s\n", st, state)
	}

	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed! You Can Close This Tab")
	ch <- client
}
