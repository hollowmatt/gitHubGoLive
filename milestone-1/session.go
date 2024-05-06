package main

type userData struct {
	Login string
	accessToken string
}

sessionStore := make(map[string]userData)