package auth_service

import (
	"net/http"

	"github.com/aidenfine/go_carmeet_backend/auth"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterAuthRoutes(r *mux.Router, db *sqlx.DB) {
	authRouter := r.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginUser(w, r, db)
	}).Methods("POST")

	authRouter.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	}).Methods("POST")

	authRouter.HandleFunc("/refresh", auth.RefreshToken).Methods("POST")
}
