package application

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitializeRouter(app *application) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_BASE_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	// public routes
	v1Router.Group(func(r chi.Router) {
		r.Post("/sign_in", app.handleUserSignIn)
		r.Post("/sign_up", app.handleUserSignUp)
		r.Get("/sign_out", app.handleUserSignOut)
		r.Get("/refresh", app.handleRefreshToken)
	})

	// private routes
	v1Router.Group(func(r chi.Router) {
		r.Use(app.authenticate)
		r.Get("/users/{user_id}", app.handleGetUserById)
	})

	router.Mount("/v1", v1Router)

	return router
}
