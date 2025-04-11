package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"Login/models"
)

var registeredUsers = []models.User{}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil || newUser.Username == "" || newUser.Password == "" {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	registeredUsers = append(registeredUsers, newUser)
	log.Printf("Usuario registrado: %s", newUser.Username)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registro exitoso"))
}
