package auth_utils

import (
	user_types "github.com/aidenfine/go_carmeet_backend/services/user/types"
	"github.com/jmoiron/sqlx"
)

func IsEmailRegistered(email string, db *sqlx.DB) bool {
	query := `SELECT id, username, email, mobile_phone, password_hash, created_at, updated_at
              FROM users WHERE email = $1`

	var user user_types.User
	err := db.Get(&user, query, email)
	return err == nil

}
