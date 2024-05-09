package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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

// create a session
func createSession(ctx context.Context, token string) (*sessionData, error) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	ghClient := github.NewClient(tc)

	u, _, err := ghClient.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	sessionId, err := getRandomString()
	if err != nil {
		return nil, err
	}
	sessionStore[sessionId] = userData{
		Login:       *u.Login,
		accessToken: token,
	}
	return &sessionData{
		ID: sessionId,
	}, nil
}

// get a session
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

func validCallback(r *http.Request) bool {

	gotState := r.URL.Query().Get("state")
	c, err := r.Cookie(oauthStateCookie)
	if err != nil {
		return false
	}
	if c.Value != gotState {
		return false
	}

	return true
}

func setCookie(writer http.ResponseWriter, name, value string, maxAge int) {
	http.SetCookie(writer, &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/",
		MaxAge: maxAge,
	})
}
