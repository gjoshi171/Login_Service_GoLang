package service

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, code int, message string, payload interface{}) {
	response, _ := json.Marshal(payload)
	//	msg, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	//w.Write(msg)
}

func ResponseWithError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
