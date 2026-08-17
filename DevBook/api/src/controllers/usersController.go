package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

//Criar usuario
func CreateUser(w http.ResponseWriter, r *http.Request)  {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil{
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.UsersModels
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
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