package server

import (
	"net/http"
)

type returnMessage struct {
	Message string `json:"message"`
}

type errorMessage struct {
	Error string `json:"error"`
}

type responceMessage struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", api)
	return mux
}
