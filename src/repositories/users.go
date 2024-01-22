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
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
