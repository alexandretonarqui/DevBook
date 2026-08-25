package repositories

import (
	"api/src/models"
	"database/sql"
)

//Representa um repositório de pubicações
type Publications struct {
	db *sql.DB
}

//Cria um repositório de publicações
func NewPublicationsRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

//Insere uma publicação do banco de dados
func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statment, erro := repository.db.Prepare(
		"insert into publications (title, content, author_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	result, erro := statment.Exec(publication.Title, publication.Content, publication.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastInsertID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertID), nil
}