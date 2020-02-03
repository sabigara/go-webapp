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
		`INSERT INTO user (id, name, email)
		 VALUES (?, ?, ?)`,
		user.ID, user.Name, user.Email,
	)
	if err != nil {
		return fmt.Errorf("mysql.user_repository.Create: %w", err)
	}
	return nil
}

func (ur *UserRepository) Get(id string) (u *api.User, err error) {
	row := ur.QueryRow(
		`SELECT id, name, email
		 FROM user
		 WHERE id = ?`,
		id,
	)
	u = &api.User{}
	err = row.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api.ErrResourceNotFound
		}
		return nil, fmt.Errorf("mysql.user_repository.Get: %w", err)
	}
	return
}
