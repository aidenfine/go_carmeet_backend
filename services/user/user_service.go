package user

import (
	"encoding/json"
	"net/http"

	user_types "github.com/aidenfine/go_carmeet_backend/services/user/types"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	// Get the email from the URL parameters
	vars := mux.Vars(r)
	email := vars["email"]

	// Define the query to get the user by email
	query := `SELECT id, username, email, mobile_phone, password, created_at, updated_at
              FROM users WHERE email = $1`

	var user user_types.User
	err := db.Get(&user, query, email)
	if err != nil {
		// If no user is found, return 404
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Encode the user as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
