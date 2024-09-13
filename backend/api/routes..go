package api

import (
	"encoding/json"
	"net/http"

	"github.com/CrossStack-Q/Vacation-App/models"
	"github.com/CrossStack-Q/Vacation-App/types"
	"github.com/CrossStack-Q/Vacation-App/utils"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	if r.Method == http.MethodPost {
		var user types.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Hash password
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword

		// Save user in DB
		if err := models.CreateUser(user); err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	if r.Method == http.MethodPost {
		var user types.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Check if user exists and validate password
		dbUser, err := models.GetUserByEmail(user.Email)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Handle login session or token (for simplicity, not implemented here)

		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	// Clear session or token (for simplicity, not implemented here)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
