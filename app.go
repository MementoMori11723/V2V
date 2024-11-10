package main

import (
	"V2V/config"
	"V2V/server"
	"log"
	"net/http"
)

func main() {
  Server := server.New()
  PORT := config.GetPort()
  mux := http.Server{
    Addr:    ":" + PORT,
    Handler: Server,
  }
  log.Println("Server started")
  log.Println(
    "live on http://localhost:" + 
    PORT,
  )
  err := mux.ListenAndServe()
  if err != nil {
    log.Println(
      "Error Starting server : ", 
      err,
    )
  }
}
