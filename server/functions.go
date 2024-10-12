package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Error(w http.ResponseWriter, err error) {
  http.Error(w, err.Error(), http.StatusInternalServerError)
}

func api(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
    Error(w, err)
		return
	}

	if len(body) == 0 {
    Error(w,errors.New("No data found"))
		return
	}

	data, err := getAiInfo(string(body))
	if err != nil {
    Error(w, err)
		return
	}

	var result responceMessage
	err = json.Unmarshal(data, &result)
	if err != nil {
    Error(w, err)
		return
	}

  content := result.Choices[0].Message.Content
  err = json.NewEncoder(w).Encode(returnMessage{
    Message: content,
  })

  if err != nil {
    Error(w, err)
    return
  }
}
