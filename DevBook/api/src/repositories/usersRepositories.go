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
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick Like ?",
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

//Busca usuário por ID no banco de dados
func (repository UsersRepo) FindByID(ID uint64) (models.UsersModels, error) {
	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.UsersModels{}, erro
	}
	defer lines.Close()

	var user models.UsersModels

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.UsersModels{}, erro
		}
	}

	return user, nil
}

//Atualiza as informações de um usuário no banco de dados
func (repository UsersRepo) UpdateUser(ID uint64, user models.UsersModels) error {
	statement, erro := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

//Exclui os dados de um usuário no banco de dados
func (repository UsersRepo) Delete(ID uint64) error {
	statement, erro := repository.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

//Busca um usuário por Email, retornando seu ID e Senha com hash
func (repository UsersRepo) FindByEmail(email string) (models.UsersModels, error) {
	line, erro := repository.db.Query(
		"select id, password from users where email = ?", email,
	)
	if erro != nil {
		return models.UsersModels{}, erro
	}
	defer line.Close()

	var user models.UsersModels

	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.UsersModels{}, erro
		}
	}

	return user, nil
}

//Permite que um usuário siga outro usuário
func (repository UsersRepo) Follow(userID, followID uint64) error {
	statment, erro := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro := statment.Exec(userID, followID); erro != nil {
		return erro
	}

	return nil
}

//Permite que um usuário pare de seguir outro usuário
func (repository UsersRepo) UnfollowUser(userID, followID uint64) error {
	statement, erro := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followID); erro != nil {
		return erro
	}

	return nil
}

//Traz todos os seguidores de um usuário
func (repository UsersRepo) SearchFollowers(userID uint64) ([]models.UsersModels, error) {
	lines, erro := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdat
		from users u inner join followers s on u.id = s.follower_id where s.user_id = ?`,
		userID,
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

//Traz todos os seguidores que um determinado usuário está seguindo
func (repository UsersRepo) SearchFollowing(userID uint64) ([]models.UsersModels, error) {
	lines, erro := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdat
		from users u inner join followers s on u.id = s.user_id where s.follower_id = ?`,
		userID,
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