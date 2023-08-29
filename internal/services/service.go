package services

import (
	"segmenatationService/internal/contracts"
	"segmenatationService/internal/repository"
)

type Service struct {
	contracts.SegmentServiceContract
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SegmentServiceContract: NewSegmentService(repo.SegmentRepositoryContract),
	}
}
