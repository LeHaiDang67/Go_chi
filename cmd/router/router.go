package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Handler returns the http handler that handles all requests
func Handler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-Id-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Route("/user", func(userRouter chi.Router) {
		userRouter.Get("/", getUser(db))
		userRouter.Post("/", addUser(db))
		userRouter.Delete("/", deleteUser(db))
	})

	return r
}
