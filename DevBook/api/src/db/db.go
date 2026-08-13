package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

//Abre a conexão com o banco de dados e a retorna
func Connection() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringDBConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}