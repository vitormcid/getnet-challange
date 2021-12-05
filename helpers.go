// app.go

package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func encryptPassword(password string) string {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
