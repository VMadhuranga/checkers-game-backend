package application

import (
	"github.com/VMadhuranga/checkers-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
)

type Application struct {
	Queries                               *database.Queries
	Validate                              *validator.Validate
	AccessTokenSecret, RefreshTokenSecret string
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
