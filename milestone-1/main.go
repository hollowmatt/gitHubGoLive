package main

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var conf = &oauth2.Config{
	ClientID:     os.Getenv("ghclient"),
	ClientSecret: os.Getenv("ghsecret"),
	Scopes:       []string{"repo", "user"},
	Endpoint: oauth2.Endpoint{
		TokenURL: "https://github.com/login/oauth/authorize",
		AuthURL:  "http://localhost:8080/github/callback",
	},
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/github/callback", githubCallbackHandler)
	http.ListenAndServe(":8080", nil)
}
