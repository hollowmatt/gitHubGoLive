package main

import (
	"crypto/rand"
	"encoding/base64"
)

type userData struct {
	Login       string
	accessToken string
}

var sessionStore = make(map[string]userData)

type sessionData struct {
	ID string
}

func getRandomString() (string, error) {
	c := 32
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// stub create a session
func createSession(token string) (string, error) {
	return token, nil
}

// stub get a session
func getSession(id string) (string, error) {
	return id, nil
}

func setCookie(value string) string {
	if value == "" {
		return "error"
	} else {
		return value
	}
}
