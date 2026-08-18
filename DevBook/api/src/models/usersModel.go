package models

import (
	"errors"
	"strings"
	"time"
)

//Representa um usuario utilizando a rede social
type UsersModels struct {
	ID uint64 			`json:"id,omitempty"`
	Name string 		`json:"name,omitempty"`
	Nick string 		`json:"nick,omitempty"`
	Email string 		`json:"email,omitempty"`
	Password string 	`json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

func (user *UsersModels) Preparer() error {
	if erro := user.validation(); erro != nil {
		return erro
	}

	user.format()
	return nil
}

//Chama os métodos pra formatar e validar os usuários recebidos
func (user *UsersModels) validation() error {
	if user.Name == "" {
		return errors.New("Campo Nome preenchimento obrigatório")
	}

	if user.Nick == "" {
		return errors.New("Campo Nick preenchimento obrigatório")
	}

	if user.Email == "" {
		return errors.New("Campo Email preenchimento obrigatório")
	}

	if user.Password == "" {
		return errors.New("Campo Password preenchimento obrigatório")
	}

	return nil
}

func (user *UsersModels) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}