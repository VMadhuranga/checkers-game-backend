package application

import (
	"net/http"
	"time"
)

func (app *application) handleUserSignOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, initializeJwtCookie(-1*time.Second, ""))
	respondWithJson(w, 204, nil)
}
