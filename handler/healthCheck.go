package handler

import (
	"fmt"
	"net/http"
	"github.com/rafaelfcads/file-api/helper"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Healthcheck success")
	helper.RespondWithJSON(w, http.StatusOK, "Healthcheck success")
}
