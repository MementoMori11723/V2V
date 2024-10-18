package server

import (
	"V2V/server/ai"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Error(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorMessage{
		Error: err.Error(),
	})
}

func api(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}

	if len(body) == 0 {
		Error(
      w, errors.New("No data found"),
      http.StatusNotFound,
    )
		return
	}

	if string(body) == " " {
		Error(
      w, errors.New("Please enter some data"),
      http.StatusNotFound,
    )
		return
	}

	data, err := ai.Talk(string(body))
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}

	var result responceMessage
	err = json.Unmarshal(data, &result)
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}

	content := result.Choices[0].Message.Content
	err = json.NewEncoder(w).Encode(returnMessage{
		Message: content,
	})

	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}
}
