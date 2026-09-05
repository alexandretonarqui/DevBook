package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// LoadLoginPage carrega a tela de Login
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

// LoadUserSubmitPage carrega a página de cadastro de usuário
func LoadUserSubmitPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "registration.html", nil)
}

// LoadHomePage carrega a página principal com as publicações
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, erro := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCodeError(w, response)
		return
	}

	var publications []models.Publication
	if erro = json.NewDecoder(response.Body).Decode(&publications); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// LoadUpdatePublicationPage carrega a página de edição da publicação
func LoadUpdatePublicationPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parameters["publicationID"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationID)
	response, erro := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatStatusCodeError(w, response)
		return
	}

	var publication models.Publication
	if erro = json.NewDecoder(response.Body).Decode(&publication); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecTemplate(w, "update-publication.html", publication)
}
