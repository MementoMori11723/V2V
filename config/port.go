package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func init() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	if PORT == "" {
		panic("PORT is not set")
	}
}

func GetPort() string {
	return PORT
}
