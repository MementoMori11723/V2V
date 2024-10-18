package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	API_KEY string
	URL     string
)

func init() {
	godotenv.Load()
	API_KEY = os.Getenv("API_KEY")
	URL = os.Getenv("URL")
	if API_KEY == "" || URL == "" {
		panic("API_KEY or URL is not set")
	}
}

func GetApiDetails() (string, string) {
	return API_KEY, URL
}
