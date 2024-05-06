package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var conf = &oauth2.Config{
	ClientID:     os.Getenv("ghclient"),
	ClientSecret: os.Getenv("ghsecret"),
	Scopes:       []string{"repo", "user"},
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func githubCallbackHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello callback, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/github/callback", githubCallbackHandler)
	http.ListenAndServe(":8080", nil)
}
