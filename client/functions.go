package client

import (
	"html/template"
	"net/http"
)

var (
	Path = "client/pages/"
)

func home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles(
		Path +
			"index.html",
	)
	if err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
	}
	if err = temp.Execute(w, nil); err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
	}
}
