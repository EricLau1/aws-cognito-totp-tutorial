package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, statusCode int, data interface{}) {
	AddHeaderWithStatus(w, "Content-Type", "application/json", statusCode)
	respond(w, data)
}

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	AddHeaderWithStatus(w, "Content-Type", "application/json", statusCode)
	respond(w, BuildJsonError(err))
}

func AddHeaderWithStatus(w http.ResponseWriter, headerKey, headerValue string, statusCode int) {
	w.Header().Set(headerKey, headerValue)
	w.WriteHeader(statusCode)
}

func respond(w http.ResponseWriter, data interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

type JsonError struct {
	Error string `json:"error"`
}

func BuildJsonError(err error) (e JsonError) {
	if err != nil {
		e.Error = err.Error()
	} else {
		e.Error = "Ocorreu um Erro!"
	}
	log.Println("Error: ", err.Error())
	return
}
