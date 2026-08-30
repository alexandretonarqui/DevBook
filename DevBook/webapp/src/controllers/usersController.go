package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/responses"
)

//CreateUser chama a API para cadastrar um novo usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"name": 	r.FormValue("name"),
		"email": 	r.FormValue("email"),
		"nick": 	r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if erro!= nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(user))
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