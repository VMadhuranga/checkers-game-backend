package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (app application) handleUserSignIn(w http.ResponseWriter, r *http.Request) {
	payload := userSignInPayload{}
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

	user, err := app.queries.GetUserByUsername(r.Context(), payload.Username)
	if err != nil {
		log.Printf("%s: %s", ErrGettingUserByUsername, err)
		respondWithValidationError(w, 401, validationErrorMessagesResponse{
			Username: []string{"Incorrect username"},
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		log.Printf("%s: %s", ErrComparingPasswords, err)
		respondWithValidationError(w, 401, validationErrorMessagesResponse{
			Password: []string{"Incorrect password"},
		})
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
