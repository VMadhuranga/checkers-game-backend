package application

import (
	"database/sql"
	"os"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
)

func InitializeApplication(db *sql.DB) application {
	return application{
		queries:            database.New(db),
		validate:           validator.New(validator.WithRequiredStructEnabled()),
		accessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		refreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
	}
}
