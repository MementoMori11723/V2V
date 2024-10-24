package apis

import (
	"V2V/config"
	"errors"
)

func GetTTSResponse(data string) ([]byte, error) {
  API_KEY, URL := config.GetTtsApiDetails()
  if URL != "" && API_KEY != "" {
    return nil, errors.New("TTS API details not found")
  }
  return nil, nil
}
