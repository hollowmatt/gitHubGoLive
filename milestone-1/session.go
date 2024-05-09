package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

type userData struct {
	Login       string
	accessToken string
}

var sessionStore = make(map[string]userData)

type sessionData struct {
	ID string
}

const (
	oauthStateCookie    = "OAuthState"
	sessionCookie       = "Session"
	sessionCookieMaxAge = 10 * 60 //10 min session
)

func getRandomString() (string, error) {
	c := 32
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// check session store for a session ID
func validSessionID(sessionID string) bool {
	_, ok := sessionStore[sessionID]
	return ok
}

// stub create a session
func createSession(token string) (string, error) {
	return token, nil
}

// stub get a session
func getSession(req *http.Request) (*sessionData, error) {
	cookie, err := req.Cookie(sessionCookie)
	if err != nil {
		return nil, err
	}

	// if we get a cookie, see if session ID is in the session store
	if !validSessionID(cookie.Value) {
		return nil, fmt.Errorf("Invalid session ID")
	}
	return &sessionData{ID: cookie.Value}, nil
}

func setCookie(writer http.ResponseWriter, name, value string, maxAge int) {
	http.SetCookie(writer, &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/",
		MaxAge: maxAge,
	})
}
