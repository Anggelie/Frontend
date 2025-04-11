package main

import (
	"log"
	"net/http"

	"Login/handlers"
	"Login/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/logout", handlers.LogoutHandler)
	mux.HandleFunc("/profile", handlers.ProfileHandler)
	mux.HandleFunc("/users", handlers.UsersHandler)

	handlerWithCORS := utils.ConfigureCORS(mux)

	log.Println("Servidor corriendo en http://localhost:3000")
	err := http.ListenAndServe(":3000", handlerWithCORS)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
