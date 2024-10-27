package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeRouter(app Application) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	v1Router := chi.NewRouter()

	v1Router.Post("/sign_up", app.handleCreateUser)

	router.Mount("/v1", v1Router)

	return router
}
