package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	GPT_API_KEY string
	GPT_URL     string
  TTS_API_KEY string
  TTS_URL string
)

func init() {
	godotenv.Load()
	GPT_API_KEY = os.Getenv("GPT_API_KEY")
	GPT_URL = os.Getenv("GPT_URL")
  TTS_API_KEY = os.Getenv("TTS_API_KEY")
  TTS_URL = os.Getenv("TTS_URL")
	if GPT_API_KEY == "" || GPT_URL == "" {
    panic("GPT : API_KEY or URL is not set")
	}
  if TTS_API_KEY == "" || TTS_URL == "" {
    panic("TTS : API_KEY or URL is not set")
  }
}

func GetGptApiDetails() (string, string) {
	return GPT_API_KEY, GPT_URL
}

func GetTtsApiDetails() (string, string) {
  return TTS_API_KEY, TTS_URL
}
