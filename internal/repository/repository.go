package repository

import (
	"github.com/jmoiron/sqlx"
	"segmenatationService/internal/contracts"
	"segmenatationService/internal/repository/postgres"
)

type Repository struct {
	contracts.SegmentRepositoryContract
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SegmentRepositoryContract: postgres.NewSegmentPostgres(db)}
}
