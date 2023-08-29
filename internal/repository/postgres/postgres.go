package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"segmenatationService/config"
)

const (
	usersTable         = "users"
	segmentsTable      = "segments"
	usersSegmentsTable = "users_segments"
)

func NewPostgresDB(config *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Password))
	if err != nil {
		return nil, err
	}

	return db, nil

}
