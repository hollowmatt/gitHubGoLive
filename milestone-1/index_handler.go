package main

import (
	"fmt"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	// Check for session HTTTP cookie in the request

	// If cookie present and it's value a valid session identifier,
	// return successful response

	// if cookie not found or the value cannot be found in the sessionStore
	// initiate GitHub authentication process
}
