package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/VMadhuranga/checkers-game-backend/internal/application"
	"github.com/VMadhuranga/checkers-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s: %s", application.ErrLoadingEnv, err)
	}

	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("%s: %s", application.ErrOpeningDb, err)
	}

	app := application.Application{
		Queries:  database.New(db),
		Validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	router := application.InitializeRouter(app)

	port := os.Getenv("PORT")
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("server listening on %v", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("%s: %s", application.ErrListeningOnServer, err)
	}
}