package controllers

import "net/http"

//Criar usuario
func CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Criando um Usuário."))
}

//Buscar todos os usuários
func FindAll(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Buscando todos os Usuários."))
}

//Buscar usuário específico
func FindById(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Buscando um Usuário."))
}

//Atualizar usuário
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Atualizando Usuário."))
}

//Deletar usuário
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Deletando Usuário."))
}