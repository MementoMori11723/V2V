package main

import (
	"V2V/config"
	"V2V/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	go func() {
		Server := server.New()
		PORT := config.GetPort()
		mux := http.Server{
			Addr:    ":" + PORT,
			Handler: Server,
		}
		log.Println("Server started at http://localhost:" + PORT)
		mux.ListenAndServe()
	}()
	fmt.Scanln()
}
