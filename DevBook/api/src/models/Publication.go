package models

import (
	"errors"
	"strings"
	"time"
)

//Representa uma publicação feita por um usuário
type Publication struct {
	ID		   uint64	 `json:"id,omitempty"`
	Title	   string	 `json:"title,omitempty"`
	Content	   string	 `json:"content,omitempty"`
	AuthorID   uint64	 `json:"authorId,omitempty"`
	AuthorNick uint64 	 `json:"authorNick,omitempty"`
	Likes 	   uint64 	 `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

//Vai chamar os métodos para validar e formatar a publicação recebida
func (publication *Publication) Prepare() error {
	if erro := publication.validate(); erro != nil {
		return erro
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if publication.Content == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}