package main

import (
	"V2V/client"
	"V2V/config"
	"V2V/server"
	"log"
	"net/http"
)

func main() {
	Client := client.New()
	Server := server.New()
  PORT := config.GetPort()
	http.Handle("/", Client)
	http.Handle("/api/", http.StripPrefix("/api", Server))
	log.Println("Server started at http://localhost:" + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
