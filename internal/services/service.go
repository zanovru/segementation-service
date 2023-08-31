package services

import (
	"segmenatationService/internal/contracts"
	"segmenatationService/internal/repository"
)

type Service struct {
	contracts.SegmentServiceContract
	contracts.UserServiceContract
	contracts.UserSegmentServiceContract
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SegmentServiceContract: NewSegmentService(repo.SegmentRepositoryContract),
		UserServiceContract: NewUserService(repo.UserRepositoryContract,
			repo.SegmentRepositoryContract,
			repo.UserSegmentRepositoryContract),
		UserSegmentServiceContract: NewUserSegmentService(repo.UserRepositoryContract,
			repo.UserSegmentRepositoryContract),
	}
}
