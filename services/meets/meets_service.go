package meets_service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aidenfine/go_carmeet_backend/auth"
	meets_types "github.com/aidenfine/go_carmeet_backend/services/meets/types"
	"github.com/jmoiron/sqlx"
)

func CreateMeet(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {

	creatorId, err := auth.ExtractIDFromToken(r)
	if err != nil {
		log.Println(creatorId, "creator_Id")
		http.Error(w, "Unauthrorized: "+err.Error(), http.StatusUnauthorized)
	}
	log.Println(creatorId, "creatorId in create meet ep")

	var newMeet meets_types.Meet
	newMeet.Creator_id = creatorId
	if err := json.NewDecoder(r.Body).Decode(&newMeet); err != nil {
		http.Error(w, "Invalid Request payload", http.StatusAccepted)
		return
	}
	query := `INSERT INTO meets (creator_id, title, description, location, meet_date, theme, meet_banner, meet_thumbnail) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING MEET_ID`
	query_err := db.QueryRow(query, creatorId, newMeet.Title, newMeet.Description, newMeet.Location, newMeet.Meet_date, newMeet.Theme, newMeet.Meet_banner, newMeet.Meet_thumbnail).Scan(&newMeet.Meet_id)
	if query_err != nil {
		log.Printf("error inserting meet", query_err)
		http.Error(w, "Failed to create meet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMeet)
}
