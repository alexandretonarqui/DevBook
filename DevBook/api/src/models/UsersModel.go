package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//Representa um usuario utilizando a rede social
type UsersModels struct {
	ID		  uint64 	`json:"id,omitempty"`
	Name	  string 	`json:"name,omitempty"`
	Nick	  string 	`json:"nick,omitempty"`
	Email	  string 	`json:"email,omitempty"`
	Password  string 	`json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

func (user *UsersModels) Preparer(step string) error {
	if erro := user.validation(step); erro != nil {
		return erro
	}

	if erro := user.format(step); erro != nil {
		return erro
	}

	return nil
}

//Chama os métodos pra formatar e validar os usuários recebidos
func (user *UsersModels) validation(step string) error {
	if user.Name == "" {
		return errors.New("Campo Nome preenchimento obrigatório")
	}

	if user.Nick == "" {
		return errors.New("Campo Nick preenchimento obrigatório")
	}

	if user.Email == "" {
		return errors.New("Campo Email preenchimento obrigatório")
	}
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("Email inserido inválido")
	}

	if step == "cadastro" && user.Password == "" {
		return errors.New("Campo Password preenchimento obrigatório")
	}

	return nil
}

func (user *UsersModels) format(step string) error{
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "cadastro" {
		passwordHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passwordHash)
	}

	return nil
}