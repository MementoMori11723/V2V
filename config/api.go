package config

import (
	"os"
)

var GPT_API_KEY, GPT_URL, TTS_API_KEY, TTS_URL string

func init() {
	GPT_API_KEY = os.Getenv("GPT_API_KEY")
  GPT_URL = "https://api.openai.com/v1/chat/completions"

	if GPT_API_KEY == "" {
		panic("GPT : API_KEY or URL is not set")
	}

	TTS_API_KEY = os.Getenv("TTS_API_KEY")
	TTS_URL = "https://texttospeech.googleapis.com/v1/text:synthesize"

	if TTS_API_KEY == "" {
		panic("TTS : API_KEY or URL is not set")
	}
}

func GetGptApiDetails() (string, string) {
	return GPT_API_KEY, GPT_URL
}

func GetTtsApiDetails() (string, string) {
	return TTS_API_KEY, TTS_URL
}
