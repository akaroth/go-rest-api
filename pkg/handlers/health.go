package handlers

import (
	"encoding/json"
	"net/http"
)

func HeathCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "I am alive"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
