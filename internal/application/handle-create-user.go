package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (app Application) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	payload := createUserPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("%s: %s", ErrDecodingPayload, err)
		respondWithError(w, 400, ErrDecodingPayload.Error())
		return
	}

	err = app.Validate.Struct(payload)
	if err != nil {
		log.Printf("%s: %s", ErrValidatingPayload, err)
		respondWithValidationError(
			w,
			422,
			generateValidationErrorMessages(err.(validator.ValidationErrors)),
		)
		return
	}

	_, err = app.Queries.GetUserByUsername(r.Context(), payload.Username)
	if err == nil {
		log.Printf("%s", ErrExistingUser)
		respondWithValidationError(w, 409, validationErrorMessagesResponse{
			Username: []string{"Username already exists"},
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		log.Printf("%s: %s", ErrHashingPassword, err)
		respondWithError(w, 422, ErrHashingPassword.Error())
		return
	}

	err = app.Queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Username: payload.Username,
		Password: string(hashedPassword),
	})

	if err != nil {
		log.Panicf("%s: %s", ErrCreatingUser, err)
		respondWithError(w, 500, ErrCreatingUser.Error())
		return
	}

	respondWithJson(w, 201, nil)
}
