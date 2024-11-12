package application

import (
	"net/http"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
)

func (app *application) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userCtx("user")).(database.User)

	respondWithJson(w, 200, userResponse{
		Id:       user.ID,
		Username: user.Username,
	})
}
