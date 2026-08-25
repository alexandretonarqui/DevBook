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

//Traz uma única publicação do banco de dados
func (repository Publications) FindByID(publicationID uint64) (models.Publication, error) {
	line, erro := repository.db.Query(`
		select p.*, u.nick from
		publications p inner join users u
		on u.id = p.author_id where p.id = ?`,
		publicationID,
	)
	if erro != nil {
		return models.Publication{}, erro
	}
	defer line.Close()

	var publication models.Publication

	if line.Next() {
		if erro = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return models.Publication{}, erro
		}
	}

	return publication, nil
}