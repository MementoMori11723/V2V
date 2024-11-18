package server

import (
	"V2V/server/middleware"
	"embed"
	"net/http"
)

var (
	//go:embed pages/*.html
	Path embed.FS

	pathName = "pages/"
	layout   = "layout.html"
	routes   = map[string]func(http.ResponseWriter, *http.Request){
		"/":         home,
		"/about":    about,
		"POST /api": api,
	}
)

func New() http.Handler {
	mux := http.NewServeMux()
	assetsPath := http.StripPrefix(
		"/assets/", http.FileServer(
			http.Dir(
				"server/" + pathName + "assets",
			),
		),
	)
	mux.Handle("/assets/", assetsPath)
	for pattern, handler := range routes {
		mux.HandleFunc(pattern, handler)
	}
	newMux := middleware.Logger(mux)
	return newMux
}
