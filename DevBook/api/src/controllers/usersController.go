package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

//Criar usuario
func CreateUser(w http.ResponseWriter, r *http.Request)  {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil{
		log.Fatal(erro)
	}

	var user models.UsersModels
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connection()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewUsersRepository(db)
	repository.Create(user)
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