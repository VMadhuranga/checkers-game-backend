package application

import (
	"log"
	"net/http"
	"time"

	"github.com/VMadhuranga/checkers-game-backend/internal/database"
)

func (app *application) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userCtx("user")).(database.User)

	err := app.queries.DeleteUserById(r.Context(), user.ID)
	if err != nil {
		log.Printf("%s: %s", ErrDeletingUserById, err)
		respondWithError(w, 400, ErrDeletingUserById.Error())
		return
	}

	http.SetCookie(w, initializeJwtCookie(-1*time.Second, ""))
	respondWithJson(w, 204, nil)
}
