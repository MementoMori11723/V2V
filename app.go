package main

import (
	"V2V/client"
	"V2V/config"
	"V2V/server"
	"net/http"
	"log"
)

func main(){
  Client := client.New()
  Server := server.New()
  http.Handle("/", Client)
  http.Handle("/api/", http.StripPrefix("/api", Server))
  log.Println("Server started at http://localhost:" + config.PORT)
  http.ListenAndServe(":"+config.PORT, nil)
}
