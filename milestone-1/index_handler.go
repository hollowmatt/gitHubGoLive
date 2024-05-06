package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	// Check for session HTTTP cookie in the request
	cookie, err := request.Cookie("Session")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			// initiate GitHub authentication process
			http.Error(writer, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(writer, "server error", http.StatusInternalServerError)
		}
		return
	}
	// If cookie present and it's value a valid session identifier,
	// return successful response
	writer.Write([]byte(cookie.Value))
}
