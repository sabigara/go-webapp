package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sabigara/go-webapp/api"
)

type UserService struct {
	*sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (us *UserService) Create(name, email string) (*api.User, error) {
	u := api.NewUser(name, email)
	_, err := us.Exec(
		`INSERT INTO user (id, name, email)
		 VALUES (?, ?, ?)`,
		u.ID, u.Name, u.Email,
	)
	if err != nil {
		return nil, fmt.Errorf("mysql Create: %w", err)
	}

	return u, nil
}

func (us *UserService) Get(id string) (u *api.User, err error) {
	row := us.QueryRow(
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
		return nil, fmt.Errorf("mysql Get: %w", err)
	}
	return
}
