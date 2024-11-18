package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
	"github.com/go-playground/validator/v10"
)

func (app *application) handleUpdateUsernameById(w http.ResponseWriter, r *http.Request) {
	payload := updateUsernamePayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("%s: %s", ErrDecodingPayload, err)
		respondWithError(w, 400, ErrDecodingPayload.Error())
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

	user := r.Context().Value(userCtx("user")).(database.User)
	u, err := app.queries.GetUserByUsername(r.Context(), payload.NewUsername)
	if err == nil && user.ID != u.ID {
		log.Printf("%s", ErrExistingUser)
		respondWithValidationError(w, 409, validationErrorMessagesResponse{
			NewUsername: []string{"New username already exists"},
		})
		return
	}

	err = app.queries.UpdateUsernameById(r.Context(), database.UpdateUsernameByIdParams{
		ID:       user.ID,
		Username: payload.NewUsername,
	})
	if err != nil {
		log.Printf("%s: %s", ErrUpdatingUsernameById, err)
		respondWithError(w, 500, ErrUpdatingUsernameById.Error())
		return
	}

	respondWithJson(w, 204, nil)
}
