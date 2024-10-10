package config

import (
  "github.com/joho/godotenv"
  "os"
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
