package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
)

//CreatePublication chama a API para cadastrar uma nova publicação no banco de dados
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, erro := json.Marshal(map[string]string{
		"title": r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, erro := requests.MakeAuthRequest(r, http.MethodPost, url, bytes.NewBuffer(publication))
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}