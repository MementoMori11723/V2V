package client

import (
	"V2V/config"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  temp, err := template.ParseFiles(config.Path + "index.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  if err = temp.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
