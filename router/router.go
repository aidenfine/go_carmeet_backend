package router

import (
	"log"
	"net/http"

	"github.com/aidenfine/go_carmeet_backend/middleware"
	auth_service "github.com/aidenfine/go_carmeet_backend/services/auth"
	meets_service "github.com/aidenfine/go_carmeet_backend/services/meets"
	user_routes "github.com/aidenfine/go_carmeet_backend/services/user"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func StartServer(db *sqlx.DB) error {
	r := mux.NewRouter()

	registerPublicRoutes(r, db)

	registerProtectedRoutes(r, db)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	}).Methods("GET")

	log.Println("Server running on port 8080...")
	return http.ListenAndServe(":8080", r)
}

func registerPublicRoutes(r *mux.Router, db *sqlx.DB) {
	auth_service.RegisterAuthRoutes(r, db)
}

func registerProtectedRoutes(r *mux.Router, db *sqlx.DB) {
	protectedRouter := r.PathPrefix("").Subrouter()
	protectedRouter.Use(middleware.JWTMiddleware)
	user_routes.RegisterUserRoutes(protectedRouter, db)
	meets_service.RegisterMeetsRoutes(r, db)
}
