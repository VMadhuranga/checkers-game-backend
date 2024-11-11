package application

import (
	"time"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type application struct {
	queries                                 *database.Queries
	validate                                *validator.Validate
	accessTokenSecret, refreshTokenSecret   string
	accessTokenExpTime, refreshTokenExpTime time.Duration
}

type createUserPayload struct {
	Username        string `json:"username,omitempty" validate:"required,alpha"`
	Password        string `json:"password,omitempty" validate:"required,alphanum,min=5"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"eqfield=Password"`
}

type userSignInPayload struct {
	Username string `json:"username,omitempty" validate:"required,alpha"`
	Password string `json:"password,omitempty" validate:"required,alphanum"`
}

type validationError struct {
	field, tag string
}

type validationErrorMessagesResponse struct {
	Username        []string `json:"username,omitempty"`
	Password        []string `json:"password,omitempty"`
	ConfirmPassword []string `json:"confirm_password,omitempty"`
}

type userResponse struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
}
