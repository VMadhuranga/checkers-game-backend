package application

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app application) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("%s: %s", ErrParsingUserIdParamToUUID, err)
		respondWithError(w, 400, ErrParsingUserIdParamToUUID.Error())
		return
	}

	user, err := app.queries.GetUserById(r.Context(), userId)
	if err != nil {
		log.Printf("%s: %s", ErrGettingUserById, err)
		respondWithError(w, 404, ErrGettingUserById.Error())
		return
	}

	respondWithJson(w, 200, userResponse{
		Id:       user.ID,
		Username: user.Username,
	})
}
