package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) handleUpdatePasswordById(w http.ResponseWriter, r *http.Request) {
	payload := updatePasswordPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("%s: %s", ErrDecodingPayload, err)
		respondWithError(w, 400, ErrDecodingPayload.Error())
		return
	}

	user := r.Context().Value(userCtx("user")).(database.User)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.OldPassword))
	if err != nil {
		log.Printf("%s: %s", ErrComparingPasswords, err)
		respondWithValidationError(w, 401, validationErrorMessagesResponse{
			OldPassword: []string{"Incorrect old password"},
		})
		return
	}

	err = app.validate.Struct(payload)
	if err != nil {
		log.Printf("%s: %s", ErrValidatingPayload, err)
		respondWithValidationError(
			w,
			422,
			generateValidationErrorMessages(err.(validator.ValidationErrors)),
		)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.NewPassword), 10)
	if err != nil {
		log.Printf("%s: %s", ErrHashingPassword, err)
		respondWithError(w, 422, ErrHashingPassword.Error())
		return
	}

	err = app.queries.UpdatePasswordById(r.Context(), database.UpdatePasswordByIdParams{
		ID:       user.ID,
		Password: string(hashedPassword),
	})
	if err != nil {
		log.Printf("%s: %s", ErrUpdatingPasswordById, err)
		respondWithError(w, 500, ErrUpdatingPasswordById.Error())
		return
	}

	respondWithJson(w, 204, nil)
}
