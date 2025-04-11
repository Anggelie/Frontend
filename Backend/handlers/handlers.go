package handlers

import (
	"encoding/json"
	"login/utils"
	"net/http/httputil"

	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.NewErrorResponse(message))
}
