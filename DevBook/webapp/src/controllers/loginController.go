package controllers

import "net/http"

//LoadLoginPage vai carregar a tela de Login
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de login"))
}