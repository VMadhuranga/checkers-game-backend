package application

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Printf("%s: %s", ErrEncodingPayload, err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload != nil {
		w.Write(response)
	}
}
