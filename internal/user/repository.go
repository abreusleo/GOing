package user

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUserByID(userID int) (string, error) {
	var name string
	query := "SELECT name FROM users WHERE id = ?"
	err := r.db.QueryRow(query, userID).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("ID %d not found", userID)
		}
		return "", fmt.Errorf("Error querying user with ID %d: %v", userID, err)
	}
	return name, nil
}
