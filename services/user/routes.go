package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(r *mux.Router, db *sqlx.DB) {
	userRouter := r.PathPrefix("/v1/users").Subrouter()

	userRouter.HandleFunc("/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		// get user by id
	}).Methods("GET")

}
