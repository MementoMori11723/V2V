package apis

import (
	"V2V/config"
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

func PostRequest(URL string, data interface{}, API_KEY string) ([]byte, error) {
  return nil, nil
}

func GetTTSResponse(data string) ([]byte, error) {
	API_KEY, URL := config.GetTtsApiDetails()
  Responce := TTSResponse{
    AudioConfig: AudioConfig{
      AudioEncoding: "LINEAR16",
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
  return PostRequest(
    URL, Responce, 
    API_KEY,
  )
}
