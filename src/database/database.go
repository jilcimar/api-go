package database

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conexão com o banco de dados
func Conect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexction)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
