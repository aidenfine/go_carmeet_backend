package meets_service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterMeetsRoutes(r *mux.Router, db *sqlx.DB) {
	meetsRouter := r.PathPrefix("/v1/meets").Subrouter()
	meetsRouter.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		CreateMeet(w, r, db)
	}).Methods("POST")
}
