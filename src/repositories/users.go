package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *Users {
	return &Users{db}
}

// Criar usuário
func (repository Users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultant, erro := statement.Exec(user.Name, user.Username, user.Email, user.Password)

	if erro != nil {
		return 0, erro
	}

	lastID, erro := resultant.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}
