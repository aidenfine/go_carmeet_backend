package auth_service

import (
	"net/http"

	"github.com/aidenfine/go_carmeet_backend/auth"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterAuthRoutes(r *mux.Router, db *sqlx.DB) {
	authRouter := r.PathPrefix("/auth").Subrouter()

	// Login
	authRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginUser(w, r, db)
	}).Methods("POST")

	// Create user
	authRouter.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	}).Methods("POST")

	// verify token
	authRouter.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		VerifyJWTToken(w, r, db)

	}).Methods("GET")

	// logout/clear jwt token
	authRouter.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		Logout(w)
	}).Methods("POST")

	// refresh token
	authRouter.HandleFunc("/refresh", auth.RefreshToken).Methods("POST")
}
