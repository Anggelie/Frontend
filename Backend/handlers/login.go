package handlers

import (
	"login/handlers"
	"login/models"
	"login/utils"
)

func PostLoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondWithError(w, http.StatusBadRequest, "Cuerpo inválido")
			return
		}

		if req.Username == "" || req.Password == "" {
			respondWithError(w, http.StatusBadRequest, "Usuario y contraseña requeridos")
			return
		}

		var storedHash string
		var userID int64
		err := db.QueryRowContext(r.Context(),
			"SELECT id, password_hash FROM users WHERE username = ?",
			req.Username,
		).Scan(&userID, &storedHash)

		if err != nil {
			if err == sql.ErrNoRows {
				respondWithError(w, http.StatusUnauthorized, "Usuario o contraseña inválidos")
			} else {
				log.Printf("Error consultando usuario '%s': %v", req.Username, err)
				respondWithError(w, http.StatusInternalServerError, "Error del servidor")
			}
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)); err != nil {
			respondWithError(w, http.StatusUnauthorized, "Usuario o contraseña inválidos")
			return
		}

		resp := models.NewSuccessResponse(models.LoginSuccessData{
			UserID:   userID,
			Username: req.Username,
		})

		respondWithJSON(w, http.StatusOK, resp)
	}
}
