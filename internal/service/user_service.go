package service

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/yelnar0112/project-management/internal/config"
	"github.com/yelnar0112/project-management/internal/domain"
)

func GetAllUsers() ([]domain.User, error) {
	rows, err := config.DB.Query("SELECT id, full_name, email, registration, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.Registration, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user *domain.User) error {
	_, err := config.DB.Exec("INSERT INTO users (id, full_name, email, registration, role) VALUES ($1, $2, $3, $4, $5)",
		user.ID, user.FullName, user.Email, user.Registration, user.Role)
	return err
}

func GetUser(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := config.DB.QueryRow("SELECT id, full_name, email, registration, role FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.FullName, &user.Email, &user.Registration, &user.Role)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &user, err
}

func UpdateUser(user *domain.User) error {
	_, err := config.DB.Exec("UPDATE users SET full_name = $1, email = $2, registration = $3, role = $4 WHERE id = $5",
		user.FullName, user.Email, user.Registration, user.Role, user.ID)
	return err
}

func DeleteUser(id uuid.UUID) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
