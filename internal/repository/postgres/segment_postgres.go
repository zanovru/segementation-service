package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"segmenatationService/internal/models"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) CreateSegment(segment *models.Segment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (slug) values ($1) RETURNING id", segmentsTable)

	row := r.db.QueryRow(query, segment.Slug)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SegmentPostgres) DeleteSegmentBySlug(slug string) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE slug = ($1) RETURNING id", segmentsTable)

	row := r.db.QueryRow(query, slug)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
