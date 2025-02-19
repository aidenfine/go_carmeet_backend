package meets_service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterMeetsRoutes(r *mux.Router, db *sqlx.DB) {
	meetsRouter := r.PathPrefix("/v1/meets").Subrouter()

	// create meet
	meetsRouter.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		CreateMeet(w, r, db)
	}).Methods("POST")

	// get meet by id
	meetsRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetMeet(w, r, db)
	}).Methods("GET")

	// get all meets by creator_id
	// TODO: add limit to this
	meetsRouter.HandleFunc("/creator/{creator_id}", func(w http.ResponseWriter, r *http.Request) {
		GetAllMeetsByCreatorId(w, r, db)
	}).Methods("GET")
}
