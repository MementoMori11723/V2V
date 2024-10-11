package server

import (
	"encoding/json"
	"V2V/config"
	"net/http"
	"bytes"
	"time"
	"io"
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

func getAiInfo(data string) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 10}
	msg := gptMessage{
		Model: "gpt-4o-mini",
		Messages: []gptType{
			{
				Role:    "user",
				Content: data,
			},
		},
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	API_KEY, URL := config.GetApiDetails()
	req, err := http.NewRequest("POST",URL,bytes.NewBuffer(jsonMsg))
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Authorization","Bearer "+API_KEY)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func api(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	if len(body) == 0 {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}
	data, err := getAiInfo(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var result responceMessage
	err = json.Unmarshal(data, &result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(result.Choices) > 0 {
		content := result.Choices[0].Message.Content
		err = json.NewEncoder(w).Encode(returnMessage{
			Message: content,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "No choices available", http.StatusInternalServerError)
		return
	}
}
