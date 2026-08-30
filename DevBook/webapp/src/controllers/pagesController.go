package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/utils"
)

//LoadLoginPage vai carregar a tela de Login
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

//LoadUserSubmitPage vai carregar a página de cadastro de usuário
func LoadUserSubmitPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "registration.html", nil)
}

//LoadHomePage carrega a página principal com as publicações
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, erro := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, erro)

	utils.ExecTemplate(w, "home.html", nil)
}