package auth_service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aidenfine/go_carmeet_backend/auth"
	"github.com/aidenfine/go_carmeet_backend/config"
	user_types "github.com/aidenfine/go_carmeet_backend/services/user/types"
	"github.com/aidenfine/go_carmeet_backend/utils"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func LoginUser(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	var user user_types.UserLoginPayload
	// Parse JSON payload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	var existingUser user_types.User
	query := `SELECT id, username, email, mobile_phone, password_hash, created_at, updated_at FROM users WHERE email = $1`
	err := db.Get(&existingUser, query, user.Email)
	log.Println(existingUser, "existing user")

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
	if !auth.VerifyPassword(user.Password, existingUser.Password_hash) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, existingUser.ID.String())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	var newUser user_types.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser.ID = uuid.New()
	hashedPassword, err := auth.HashPassword(newUser.Password_hash)
	fmt.Print(hashedPassword, "hashed password")
	if err != nil {
		http.Error(w, "Internal Server Error When created the account (Most likely from hash)", http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (id, username, email, mobile_phone, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	query_err := db.QueryRow(query, newUser.ID, newUser.Username, newUser.Email, newUser.Mobile_phone, hashedPassword).Scan(&newUser.ID)
	if query_err != nil {
		log.Panicln(query_err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := user_types.Response{
		Message: "User created successfully",
		Status:  "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
