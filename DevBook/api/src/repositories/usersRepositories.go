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
func (u UsersRepo) Create(user models.UsersModels) (uint64, error) {
	return 0, nil
}