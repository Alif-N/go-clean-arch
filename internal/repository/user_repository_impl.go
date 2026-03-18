package repository

import (
	"database/sql"
	"go-clean-architecture/internal/models"
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	query := "INSERT INTO users(name) VALUES(?)"

	result, err := r.db.Exec(query, user.Name)
	if err != nil {
		return model.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.User{}, err
	}

	user.ID = int(id)

	return user, nil
}