package application

import (
	"log"
	"net/http"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
)

func (app application) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userCtx("user")).(database.User)

	err := app.queries.DeleteUserById(r.Context(), user.ID)
	if err != nil {
		log.Printf("%s: %s", ErrDeletingUserById, err)
		respondWithError(w, 400, ErrDeletingUserById.Error())
		return
	}

	respondWithJson(w, 204, nil)
}
