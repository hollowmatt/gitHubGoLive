package main

import (
	"fmt"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	// Check for session HTTTP cookie in the request
	cookie, cErr := getSession(request)
	// If no cookie, initiate GitHub Auth
	if cErr != nil {
		//create a session
		stateToken, sErr := getRandomString()
		if sErr != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
		}
		githubLoginUrl := oauthConf.AuthCodeURL(stateToken)
		setCookie(writer, oauthStateCookie, stateToken, 600)
		http.Redirect(writer, request, githubLoginUrl, http.StatusTemporaryRedirect)
		return
	}
	// return successful response
	fmt.Fprintf(writer, "Successfully authorized to access GitHub on your behalf: %v", sessionStore[cookie.ID].Login)

}
