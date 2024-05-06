package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var oauthConf *oauth2.Config

func initOauthConfig() {

	if len(os.Getenv("ghclient")) == 0 || len(os.Getenv("ghsecret")) == 0 {
		log.Fatal("Must specific GitHub client and secret in OS environment variables")
	}

	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("ghclient"),
		ClientSecret: os.Getenv("ghsecret"),
		Scopes:       []string{"repo", "user"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://github.com/login/oauth/authorize",
			AuthURL:  "http://localhost:8080/github/callback",
		},
	}
	log.Printf("Client: %s, Secret: %s ", oauthConf.ClientID, oauthConf.ClientSecret)
}

func main() {
	//setup Oath Configuration from GitHub
	initOauthConfig()

	//handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/github/callback", githubCallbackHandler)

	//start web server
	http.ListenAndServe(":8080", nil)
}
