package services

import (
	"errors"
	"fmt"
	"segmenatationService/internal/contracts"
	"segmenatationService/internal/models"
	"segmenatationService/internal/repository/postgres"
)

type SegmentService struct {
	repo contracts.SegmentRepositoryContract
}

func NewSegmentService(repo contracts.SegmentRepositoryContract) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment *models.Segment) (int, error) {
	id, err := s.repo.CreateSegment(segment)
	if err != nil {
		if errors.Is(err, postgres.ErrUniqueConstraint) {
			return 0, fmt.Errorf("you have already added this segment")
		}
		return 0, err
	}
	return id, nil
}

func (s *SegmentService) DeleteSegmentBySlug(slug string) (int, error) {
	return s.repo.DeleteSegmentBySlug(slug)
}

func (s *SegmentService) GetSegmentsByUserId(id int) ([]models.Segment, error) {
	return s.repo.GetSegmentsByUserId(id)
}
