package client

import(
  "net/http"
)

var(
  routes = map[string]func(http.ResponseWriter, *http.Request){
    "/": home,
    "/about": about,
    "/contact": contact,
  }
)

func New() *http.ServeMux {
  mux := http.NewServeMux()
  for pattern, handler := range routes {
    mux.HandleFunc(pattern, handler)
  }
  return mux
}
