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
}

func GetGPTResponce(data string) ([]byte, error) {
  log.Println("In GPT function")
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	msg := gptMessage{
		Model: "gpt-4o-mini",
		Messages: []gptType{
			{
				Role:    "system",
				Content: "Your name is Echo Flow, and you are a chatbot. You are designed to help users with their queries. How can I help you today?, also you can only respond in paragraph style text only, you can't respond in markdown or html, and if the user asks you to translate something, you should act like a translator and translate the text to the desired language and use latest data of year : " + strconv.Itoa(time.Now().Year()),
			},
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
  log.Println(string(body))
	if err != nil {
		return nil, err
	}

	return body, nil
}
