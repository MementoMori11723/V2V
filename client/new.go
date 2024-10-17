package client

import (
	"net/http"
)

var (
	routes = map[string]func(http.ResponseWriter, *http.Request){
		"/": home,
	}
	staticDir = "client/pages/static"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	for pattern, handler := range routes {
		mux.HandleFunc(pattern, handler)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	return mux
}
