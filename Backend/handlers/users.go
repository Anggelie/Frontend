package handlers

import (
	"encoding/json"
	"net/http"
)

type PublicUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

var allUsers = []PublicUser{
	{ID: 1, Username: "Anggelie"},
	{ID: 2, Username: "Dennis"},
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}
