package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akaroth/go-rest-api/config"
	"github.com/akaroth/go-rest-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusBadRequest)
		return
	}
	user.Password = string(hashedPassword)
	config.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
