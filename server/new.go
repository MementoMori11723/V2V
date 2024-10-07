package server

import(
  "net/http"
)

func New() *http.ServeMux {
  mux := http.NewServeMux()
  mux.HandleFunc("POST /", api)
  return mux
}
