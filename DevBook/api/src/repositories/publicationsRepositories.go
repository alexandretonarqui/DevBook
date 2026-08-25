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

//Traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repository Publications) FindAllPublications(userID uint64) ([]models.Publication, error) {
	lines, erro := repository.db.Query(`
	select distinct p.*, u.nick from publications p
	inner join users u on u.id = p.author_id
	inner join followers s on p.author_id = s.user_id
	where u.id = ? or s.follower_id = ?
	order by 1 desc`,
		userID, userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if erro = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

//Altera os dados de uma publicação no banco de dados
func (repository Publications) UpdatePublication(publicationID uint64, publication models.Publication) error {
	statment, erro := repository.db.Prepare("update publications set title = ?, content = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(publication.Title, publication.Content, publicationID); erro != nil {
		return erro
	}

	return nil
}

//Exclui uma publicação do banco de dados
func (repository Publications) DeletePublication(publicationID uint64) error {
	statment, erro := repository.db.Prepare("delete from publications where id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}