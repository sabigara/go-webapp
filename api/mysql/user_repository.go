package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sabigara/go-webapp/api"
)

type UserRepository struct {
	*sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Save(user *api.User) error {
	_, err := ur.Exec(
		`INSERT INTO user (id, name, email, password)
		 VALUES (?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE name = ?
		 `,
		user.ID, user.Name, user.Email, user.Password,
		user.Name,
	)
	if err != nil {
		return fmt.Errorf("mysql.user_repository.Create: %w", err)
	}
	return nil
}

func (ur *UserRepository) GetById(id string) (u *api.User, err error) {
	row := ur.QueryRow(
		`SELECT id, name, email, password
		 FROM user
		 WHERE id = ?`, id,
	)
	u = &api.User{}
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("mysql.user_repository.Get: %w", api.ErrResourceNotFound)
		}
		return nil, fmt.Errorf("mysql.user_repository.Get: %w", err)
	}
	return
}

func (ur *UserRepository) GetByEmail(email string) (u *api.User, err error) {
	row := ur.QueryRow(
		`SELECT id, name, email, password
		 FROM user
		 WHERE email = ?`, email,
	)
	u = &api.User{}
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("mysql.user_repository.Get: %w", api.ErrResourceNotFound)
		}
		return nil, fmt.Errorf("mysql.user_repository.Get: %w", err)
	}
	return
}
