package config

import (
	"log"
	"os"
)

var (
	PORT                          string
	GPT_URL, TTS_API_KEY, TTS_URL string
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Printf("PORT is not set. Defaulting to 8080")
		PORT = "8080"
	}
	GPT_URL = "http://echo_ai:11434/api"
	TTS_API_KEY = os.Getenv("TTS_API_KEY")
	TTS_URL = "https://texttospeech.googleapis.com/v1/text:synthesize"
	if TTS_API_KEY == "" {
		panic("TTS : API_KEY or URL is not set")
	}
}

func GetPort() string {
	return PORT
}

func GetGptApiDetails() string {
	return GPT_URL
}

func GetTtsApiDetails() (string, string) {
	return TTS_API_KEY, TTS_URL
}
