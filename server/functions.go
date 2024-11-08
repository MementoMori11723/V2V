package server

import (
	"V2V/server/apis"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type returnMessage struct {
	Message      string `json:"message"`
	AudioMessage string `json:"audioMessage"`
}

type audioMessage struct {
	AudioContent string `json:"audioContent"`
}

type responceMessage struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Default error handler (returns JSON)
func Error(w http.ResponseWriter, err error, statusCode int) {
	type errorMessage struct {
		Error string `json:"error"`
	}

	w.WriteHeader(statusCode)
	log.Println("Error : ", err)

	json.NewEncoder(w).Encode(errorMessage{
		Error: err.Error(),
	})
}

// Default error handler (returns HTML)
func HTMLError(w http.ResponseWriter, Err error, statusCode int) {
	type errorMessage struct {
		Title        string
		ErrorMessage string
		StatusCode   int
	}

	temp, err := template.ParseFiles(
		Path + "error.html",
	)

	if err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
	}

	w.WriteHeader(statusCode)

  data := errorMessage{
		Title: fmt.Sprintf("%d Error", statusCode),
		ErrorMessage: Err.Error(),
		StatusCode:   statusCode,
	}

	if err = temp.Execute(w, data); err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
	}
}

// this is /auth endpoint
func auth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth"))
}

// this is / endpoint
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HTMLError(w, errors.New("404 page not found"), http.StatusNotFound)
		return
	}

	temp, err := template.ParseFiles(
		Path +
			"index.html",
	)

	if err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
	}

	if err = temp.Execute(w, nil); err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
	}
}

// this is /api endpoint
func api(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}

	if len(body) == 0 {
		Error(
			w, errors.New("please enter some data"),
			http.StatusNotFound,
		)
		return
	}

	data, err := apis.GetGPTResponce(string(body))
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
	audioFile, err := apis.GetTTSResponse(content)
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}

	var audio audioMessage
	err = json.Unmarshal(audioFile, &audio)
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(returnMessage{
		Message:      content,
		AudioMessage: "data:audio/mp3;base64," + audio.AudioContent,
	})
	if err != nil {
		Error(w, err, http.StatusNotFound)
		return
	}
}
