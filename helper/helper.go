package helper

import (
	base64 "encoding/base64"
	"encoding/json"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithBase64(w http.ResponseWriter, code int, payload *aws.WriteAtBuffer) {
	sEnc := base64.StdEncoding.EncodeToString([]byte(payload.Bytes()))
	response, _ := json.Marshal(sEnc)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithBuffer(w http.ResponseWriter, code int, payload *aws.WriteAtBuffer) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(code)
	w.Write([]byte(payload.Bytes()))
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
