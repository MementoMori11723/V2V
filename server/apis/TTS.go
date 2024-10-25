package apis

import (
	"V2V/config"
	"errors"
)

type AudioConfig struct {
	AudioEncoding    string   `json:"audioEncoding"`
	EffectsProfileId []string `json:"effectsProfileId"`
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
	API_KEY, URL := config.GetTtsApiDetails()
	if URL != "" && API_KEY != "" {
		return nil, errors.New("TTS API details not found")
	}
	return nil, nil
}
