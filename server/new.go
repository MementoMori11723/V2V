package server

import (
	"V2V/server/middleware"
	"net/http"
)

var (
	Path   = "server/pages/"
	routes = map[string]func(http.ResponseWriter, *http.Request){
		"/":         home,
		"/auth":     auth,
		"POST /api": api,
	}
)

func New() http.Handler {
	mux := http.NewServeMux()
	assetsPath := http.StripPrefix("/assets/", http.FileServer(http.Dir(Path+"assets")))
	mux.Handle("/assets/", assetsPath)
	for pattern, handler := range routes {
		mux.HandleFunc(pattern, handler)
	}
	newMux := middleware.Logger(mux)
	return newMux
}
