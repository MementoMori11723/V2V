package server

import (
	"encoding/json"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode("{message: 'Hello World'}")
}
