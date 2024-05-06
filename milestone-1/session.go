package main

type userData struct {
	Login       string
	accessToken string
}

var sessionStore = make(map[string]userData)
