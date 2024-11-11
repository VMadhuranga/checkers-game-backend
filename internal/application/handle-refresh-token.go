package application

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (app *application) handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	jwtCookie, err := r.Cookie("jwt")
	if err != nil {
		log.Printf("%s: %s", ErrGettingJwtCookie, err)
		respondWithError(w, 401, ErrGettingJwtCookie.Error())
		return
	}

	jwtSub, err := validateJwt(jwtCookie.Value, app.refreshTokenSecret)
	if err != nil {
		log.Printf("%s: %s", ErrValidatingJwt, err)
		respondWithError(w, 403, ErrValidatingJwt.Error())
		return
	}

	userId, err := uuid.Parse(jwtSub)
	if err != nil {
		log.Printf("%s: %s", ErrParsingJwtSubToUUID, err)
		respondWithError(w, 400, ErrParsingJwtSubToUUID.Error())
		return
	}

	user, err := app.queries.GetUserById(r.Context(), userId)
	if err != nil {
		log.Printf("%s: %s", ErrGettingUserById, err)
		respondWithError(w, 404, ErrGettingUserById.Error())
		return
	}

	accessToken, err := createJWT(app.accessTokenExpTime, user.ID.String(), app.accessTokenSecret)
	if err != nil {
		log.Printf("%s: %s", ErrCreatingAccessToken, err)
		respondWithError(w, 400, ErrCreatingAccessToken.Error())
		return
	}

	refreshToken, err := createJWT(app.refreshTokenExpTime, user.ID.String(), app.refreshTokenSecret)
	if err != nil {
		log.Printf("%s: %s", ErrCreatingRefreshToken, err)
		respondWithError(w, 400, ErrCreatingRefreshToken.Error())
		return
	}

	http.SetCookie(w, initializeJwtCookie(app.refreshTokenExpTime, refreshToken))

	respondWithJson(
		w,
		200,
		map[string]string{"userId": user.ID.String(), "accessToken": accessToken},
	)
}
