package server

import (
	"V2V/server/apis"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
)

type returnMessage struct {
	Message      string `json:"message"`
	AudioMessage string `json:"audioMessage"`
}

type errorMessage struct {
	Error string `json:"error"`
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

func auth(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("auth"))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
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

func Error(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
  log.Println("Error : ", err)
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
