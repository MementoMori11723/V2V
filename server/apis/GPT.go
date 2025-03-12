package apis

import (
	"V2V/config"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type gptType struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type gptMessage struct {
	Model    string    `json:"model"`
	Messages []gptType `json:"messages"`
	Stream   bool      `json:"stream"`
	Raw      bool      `json:"raw"`
}

func GetGPTResponce(data string) ([]byte, error) {
	log.Println("In GPT function")
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	msg := gptMessage{
		Model: "smollm:135m",
		Messages: []gptType{
			{
				Role:    "system",
				Content: "Your name is Echo Flow, and you are a chatbot. You are designed to help users with their queries. How can I help you today?, also you can only respond in paragraph style text only, you can't respond in markdown or html, and if the user asks you to translate something, you should act like a translator and translate the text to the desired language.",
			},
			{
				Role:    "user",
				Content: data + " (only in plaintext formate)",
			},
		},
		Stream: false,
    Raw: false,
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	URL := config.GetGptApiDetails()

	req, err := http.NewRequest(
		"POST", URL+"/chat",
		bytes.NewBuffer(jsonMsg),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	log.Println(string(body))
	if err != nil {
		return nil, err
	}

	return body, nil
}
