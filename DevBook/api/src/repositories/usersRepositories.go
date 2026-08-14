package repositories

import (
	"api/src/models"
	"database/sql"
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