package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"segmenatationService/internal/models"
)

var (
	ErrSegmentNotFound = errors.New("segment not found")
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (s *SegmentPostgres) CreateSegment(segment *models.Segment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (slug) values ($1) RETURNING id", segmentsTable)

	row := s.db.QueryRow(query, segment.Slug)
	if err := row.Scan(&id); err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == pgUniqueConstraintCode {
			return 0, ErrUniqueConstraint
		}
		return 0, err
	}
	return id, nil
}

func (s *SegmentPostgres) GetSegmentBySlug(slug string) (*models.Segment, error) {
	segment := &models.Segment{}
	query := fmt.Sprintf("SELECT id, slug FROM %s WHERE slug=$1", segmentsTable)
	if err := s.db.QueryRowx(query, slug).StructScan(segment); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSegmentNotFound
		}
		return nil, err
	}
	return segment, nil
}

func (s *SegmentPostgres) DeleteSegmentBySlug(slug string) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE slug = ($1) RETURNING id", segmentsTable)

	row := s.db.QueryRow(query, slug)
	var id int
	if err := row.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrSegmentNotFound
		}
		return 0, err
	}
	return id, nil
}

func (s *SegmentPostgres) GetSegmentsByUserId(id int) ([]models.Segment, error) {
	var segments []models.Segment
	query := fmt.Sprintf("SELECT s.id AS id, s.slug AS slug FROM %s AS us join %s AS s ON us.segment_id = s.id WHERE us.user_id = ($1)", usersSegmentsTable, segmentsTable)
	if err := s.db.Select(&segments, query, id); err != nil {
		return nil, err
	}
	return segments, nil
}
