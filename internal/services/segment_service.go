package services

import (
	"segmenatationService/internal/contracts"
	"segmenatationService/internal/models"
)

type SegmentService struct {
	repo contracts.SegmentRepositoryContract
}

func NewSegmentService(repo contracts.SegmentRepositoryContract) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s SegmentService) CreateSegment(segment *models.Segment) (int, error) {
	return s.repo.CreateSegment(segment)
}

func (s SegmentService) DeleteSegmentBySlug(slug string) (int, error) {
	return s.repo.DeleteSegmentBySlug(slug)
}
