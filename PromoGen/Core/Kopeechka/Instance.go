package Kopeechka

import "net/http"

type Instance struct {
	Client *http.Client
	Email  string
	ID     string
	ApiKey string
}
