package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rafaelfcads/file-api/helper"
	"github.com/rafaelfcads/file-api/model"
)

func Document(w http.ResponseWriter, r *http.Request) {
	var documentFile model.DocumentFile

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if err := r.Body.Close(); err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if err := json.Unmarshal(body, &documentFile); err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	buffer, err := helper.JsonToXlsx(documentFile)

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	location, err := helper.PublishToS3(documentFile.FileName, buffer)

	fmt.Println("S3 file location:", location)
	helper.RespondWithJSON(w, http.StatusOK, location)
}
