package user

import (
	"encoding/json"
	"log"
	"net/http"

	user_types "github.com/aidenfine/go_carmeet_backend/services/user/types"
	"github.com/aidenfine/go_carmeet_backend/utils"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	vars := mux.Vars(r)
	email := vars["email"]

	query := `SELECT id, username, email, mobile_phone, password_hash, created_at, updated_at
              FROM users WHERE email = $1`

	var user user_types.User
	err := db.Get(&user, query, email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
func GetUserByID(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Printf("Looking for user with ID: %s", id)

	var user user_types.UserResponsePayload
	query := `SELECT id, username, email, mobile_phone, created_at, updated_at FROM users WHERE id = $1`
	err := db.Get(&user, query, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
