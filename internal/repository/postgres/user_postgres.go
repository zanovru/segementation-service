package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"segmenatationService/internal/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CreateUser(user *models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email) values ($1) RETURNING id", usersTable)

	row := u.db.QueryRow(query, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserPostgres) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1", usersTable)
	if err := u.db.QueryRowx(query, email).StructScan(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *UserPostgres) GetUser(id int) (*models.User, error) {
	user := &models.User{}
	query := fmt.Sprintf("SELECT id, email FROM %s ", usersTable)
	if err := u.db.QueryRowx(query, id).StructScan(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *UserPostgres) GetUsers() ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf("SELECT id, email FROM %s", usersTable)
	if err := u.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}
