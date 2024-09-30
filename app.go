package main

import (
	"V2V/client"
	"V2V/server"
	"fmt"
	"net/http"
)

func main(){
  Client := client.New()
  Server := server.New()
  http.Handle("/", Client)
  http.Handle("/api/", http.StripPrefix("/api", Server))
  fmt.Println("Server started at http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}
