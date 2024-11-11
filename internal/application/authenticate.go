package application

import (
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			log.Printf("%s: %s", ErrValidatingBearerToken, "missing bearer token")
			respondWithError(w, 401, ErrValidatingBearerToken.Error())
			return
		}
		if !strings.HasPrefix(bearerToken, "Bearer ") {
			log.Printf("%s: %s", ErrValidatingBearerToken, "missing 'Bearer' prefix")
			respondWithError(w, 401, ErrValidatingBearerToken.Error())
			return
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")
		if token == "" {
			log.Printf("%s: %s", ErrValidatingBearerToken, "missing token value")
			respondWithError(w, 401, ErrValidatingBearerToken.Error())
		}

		jwtSub, err := validateJwt(token, app.accessTokenSecret)
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

		_, err = app.queries.GetUserById(r.Context(), userId)
		if err != nil {
			log.Printf("%s: %s", ErrGettingUserById, err)
			respondWithError(w, 404, ErrGettingUserById.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
