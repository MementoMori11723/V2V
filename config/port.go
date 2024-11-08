package config

import (
	"log"
	"os"
)

var (
	PORT string
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
	  log.Printf("PORT is not set. Defaulting to 8080")
    PORT = "8080"
	}
}

func GetPort() string {
	return PORT
}
