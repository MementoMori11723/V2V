package client

import(
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Home"))
}

func about(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("About"))
}

func contact(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Contact"))
}
