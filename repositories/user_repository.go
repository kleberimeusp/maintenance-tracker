package repositories

import (
	"maintenance-tracker/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user models.User) error {
	_, err := r.db.NamedExec(`INSERT INTO users (username, password, role) VALUES (:username, :password, :role)`, &user)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE username = ?`, username)
	return user, err
}

func (r *UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE id = ?`, id)
	return user, err
}
