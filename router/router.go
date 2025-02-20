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
	"github.com/rs/cors"
)

func StartServer(db *sqlx.DB) error {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173", // Frontend URL
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type", "Authorization",
		},
		AllowCredentials: true,
	})

	registerPublicRoutes(r, db)
	registerProtectedRoutes(r, db)

	r.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // TODO: change later
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}
	}).Methods("OPTIONS")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	}).Methods("GET")

	handler := c.Handler(r)

	log.Println("Server running on port 8080...")
	return http.ListenAndServe(":8080", handler)
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
