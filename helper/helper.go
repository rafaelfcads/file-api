package helper

import (
	base64 "encoding/base64"
	"encoding/json"
	"net/http"
	"bytes"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithBase64(w http.ResponseWriter, code int, payload *bytes.Buffer) {
	sEnc := base64.StdEncoding.EncodeToString([]byte(payload.Bytes()))
	response, _ := json.Marshal(sEnc)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithInt64(w http.ResponseWriter, code int, payload int64) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
