package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectToDatabase() (*sqlx.DB, error) {
	var err error

	DB, err = sqlx.Connect("postgres", "user=postgres dbname=go_carmeet sslmode=disable password=root host=localhost")
	if err != nil {
		log.Println("Database connection error:", err)
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		log.Println("Database ping failed", err)
		return nil, err
	}
	log.Println("connected to postgres")
	return DB, nil
}
