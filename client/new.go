package client

import (
	"net/http"
)

var (
	staticDir = "client/pages/static"
	routes    = map[string]func(http.ResponseWriter, *http.Request){
		"/": home,
	}
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	for pattern, handler := range routes {
		mux.HandleFunc(pattern, handler)
	}
	return mux
}
