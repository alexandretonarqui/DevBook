package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//LoadLoginPage vai carregar a tela de Login
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}