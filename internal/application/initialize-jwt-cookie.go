package application

import (
	"net/http"
	"time"
)

func initializeJwtCookie(expiresAt time.Duration, token string) *http.Cookie {
	return &http.Cookie{
		Name:        "jwt",
		Value:       token,
		MaxAge:      int(expiresAt.Seconds()),
		Secure:      true,
		HttpOnly:    true,
		SameSite:    http.SameSiteNoneMode,
		Partitioned: true,
	}
}
