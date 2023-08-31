package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"segmenatationService/internal/models"
)

const pgUniqueConstraintCode = "23505"

var (
	ErrUniqueConstraint    = errors.New("error unique constraint")
	ErrUserSegmentNotFound = errors.New("userSegment not found")
)

type UserSegmentPostgres struct {
	db *sqlx.DB
}

func NewUserSegmentPostgres(db *sqlx.DB) *UserSegmentPostgres {
	return &UserSegmentPostgres{db: db}
}

func (u *UserSegmentPostgres) CreateUserSegment(userSegment *models.UserSegment) (int, int, error) {
	var userId, segmentId int
	query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2) RETURNING user_id, segment_id", usersSegmentsTable)
	row := u.db.QueryRow(query, userSegment.UserId, userSegment.SegmentId)
	if err := row.Scan(&userId, &segmentId); err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == pgUniqueConstraintCode {
			return 0, 0, ErrUniqueConstraint
		}

		return 0, 0, err
	}
	return userId, segmentId, nil
}

func (u *UserSegmentPostgres) DeleteUserSegment(userSegment *models.UserSegment) (int, int, error) {
	var userId, segmentId int
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_id = $2 RETURNING user_id, segment_id", usersSegmentsTable)
	row := u.db.QueryRow(query, userSegment.UserId, userSegment.SegmentId)
	if err := row.Scan(&userId, &segmentId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, 0, ErrUserSegmentNotFound
		}
		return 0, 0, err
	}
	return userId, segmentId, nil
}
