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

func getRandomString() (string, error) {
	c := 32
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
