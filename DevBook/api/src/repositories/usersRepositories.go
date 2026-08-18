package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

//Repositório de usuários
type UsersRepo struct {
	db *sql.DB
}

//Cria um repositório de usuários
func NewUsersRepository(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

//Insere um usuário no banco de dados
func (repository UsersRepo) Create(user models.UsersModels) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIDInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInsert), nil
}

//Busca todos os usuários que atendem o filtro de Name ou Nick
func (repository UsersRepo) Search(nameOrNick string) ([]models.UsersModels, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdat from users where name LIKE ? or nick Like ?",
		nameOrNick, nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.UsersModels

	for lines.Next() {
		var user models.UsersModels

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}