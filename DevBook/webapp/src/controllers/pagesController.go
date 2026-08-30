package controllers

import (
	"net/http"
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
	utils.ExecTemplate(w, "home.html", nil)
}