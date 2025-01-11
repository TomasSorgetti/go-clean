package repository

import (
	"database/sql"

	"github.com/TomasSorgetti/go-clean/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id int64) (*entity.User, error)
	FindAll() ([]*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	return err
}

func (r *userRepository) FindByID(id int64) (*entity.User, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	user := &entity.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]*entity.User, error) {
	query := "SELECT id, username, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
