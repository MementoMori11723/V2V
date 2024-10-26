package apis

import (
	"V2V/config"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type AudioConfig struct {
	AudioEncoding    string   `json:"audioEncoding"`
	Pitch            float64  `json:"pitch"`
	SpeakingRate     float64  `json:"speakingRate"`
}

type Input struct {
	Text string `json:"text"`
}

type Voice struct {
	LanguageCode string `json:"languageCode"`
	Name         string `json:"name"`
}

type TTSResponse struct {
	AudioConfig AudioConfig `json:"audioConfig"`
	Input       Input       `json:"input"`
	Voice       Voice       `json:"voice"`
}

func GetTTSResponse(data string) ([]byte, error) {
  client := &http.Client{
    Timeout: time.Second * 20,
  }
  
  Responce := TTSResponse{
    AudioConfig: AudioConfig{
      AudioEncoding: "MP3",
      Pitch: 0,
      SpeakingRate: 1,
    },
    Input: Input{
      Text: data,
    },
    Voice: Voice{
      LanguageCode: "en-US",
      Name: "en-US-Standard-G",
    },
  }

	API_KEY, URL := config.GetTtsApiDetails()

  jsonMsg, err := json.Marshal(Responce)
  if err != nil {
    return nil, err
  }

  req, err := http.NewRequest(
    "POST", URL + "?key=" + API_KEY,
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
  if err != nil {
    return nil, err
  }

  return body, nil
}
