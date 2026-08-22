package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if erro = user.Preparer("cadastro"); erro != nil {
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
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := db.Connection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, erro := repository.Search(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

//Buscar usuário específico
func FindByID(w http.ResponseWriter, r *http.Request)  {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
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
	user, erro := repository.FindByID(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

//Atualizar usuário
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro :=  authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu"))
		return
	}

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

	if erro = user.Preparer("edition"); erro != nil {
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
	if erro = repository.UpdateUser(userID, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//Deletar usuário
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o seu"))
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if erro = repository.Delete(userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//Permite que um usuário siga outro
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userID"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir a si mesmo"))
		return
	}

	db, erro := db.Connection()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if erro = repository.Follow(userID, followID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}