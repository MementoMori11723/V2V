package apis

import (
	"V2V/config"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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
		Model: "tinyllama",
		Messages: []gptType{
			{
				Role:    "system",
				Content: "Your name is Echo Flow, and you are a chatbot designed to assist users with their queries efficiently. You must always respond in a single continuous paragraph without using markdown, HTML, bullet points, headings, or special formatting. When a user requests a translation, act as a translator and provide the translated text in the desired language. Additionally, ensure that your responses reflect the most up to-date information available for the current year: " + strconv.Itoa(time.Now().Year()) + ".",
			},
			{
				Role:    "user",
				Content: data + " (only in plaintext formate)",
			},
		},
		Stream: false,
    Raw: true,
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
