package server

import(
  "net/http"
)

type returnMessage struct {
	Message string `json:"message"`
}

type responceMessage struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type gptType struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type gptMessage struct {
	Model    string    `json:"model"`
	Messages []gptType `json:"messages"`
}

func New() *http.ServeMux {
  mux := http.NewServeMux()
  mux.HandleFunc("POST /", api)
  return mux
}
