package apis

import (
	"V2V/config"
	"bytes"
	"encoding/json"
	"io"
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
}

func GetGPTResponce(data string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	msg := gptMessage{
		Model: "gpt-4o-mini",
		Messages: []gptType{
			{
				Role:    "user",
				Content: data + " in paragraph style",
			},
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	API_KEY, URL := config.GetGptApiDetails()

	req, err := http.NewRequest(
		"POST", URL,
		bytes.NewBuffer(jsonMsg),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+API_KEY,
	)

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
