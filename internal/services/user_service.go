package services

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"segmenatationService/internal/contracts"
	"segmenatationService/internal/models"
	"segmenatationService/internal/repository/postgres"
)

type UserService struct {
	userRepo        contracts.UserRepositoryContract
	segmentRepo     contracts.SegmentRepositoryContract
	userSegmentRepo contracts.UserSegmentRepositoryContract
}

func NewUserService(userRepo contracts.UserRepositoryContract,
	segmentRepo contracts.SegmentRepositoryContract,
	userSegmentRepo contracts.UserSegmentRepositoryContract) *UserService {
	return &UserService{userRepo: userRepo, segmentRepo: segmentRepo, userSegmentRepo: userSegmentRepo}
}

func (u *UserService) GetUser(id int) (*models.User, error) {
	return u.userRepo.GetUser(id)
}

func (u *UserService) CreateUserWithSegments(user *models.User) (int, error) {
	var userId int
	userFromDB, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		userId, err = u.userRepo.CreateUser(user)
		if err != nil {
			return 0, err
		}
	} else {
		userId = userFromDB.Id
	}

	for _, segmentSlugToAdd := range user.SegmentsToAdd {
		segment, err := u.segmentRepo.GetSegmentBySlug(segmentSlugToAdd)
		if err != nil {
			return 0, fmt.Errorf("can't find segment with slug: %s", segmentSlugToAdd)
		}
		userSegment := &models.UserSegment{
			UserId:    userId,
			SegmentId: segment.Id,
		}
		userId, segmentId, err := u.userSegmentRepo.CreateUserSegment(userSegment)
		if err != nil {
			if errors.Is(err, postgres.ErrUniqueConstraint) {
				return 0, fmt.Errorf("you have already added this user to slug: %s", segmentSlugToAdd)
			}
			return 0, err
		}
		log.Infof("Added user_id: %d to segment with id: %d and slug: %s", userId, segmentId, segmentSlugToAdd)
	}

	for _, segmentSlugToDelete := range user.SegmentsToDelete {
		segment, err := u.segmentRepo.GetSegmentBySlug(segmentSlugToDelete)
		if err != nil {
			return 0, fmt.Errorf("can't find segment with slug: %s", segmentSlugToDelete)
		}
		userSegment := &models.UserSegment{
			UserId:    userId,
			SegmentId: segment.Id,
		}
		userId, segmentId, err := u.userSegmentRepo.DeleteUserSegment(userSegment)
		if err != nil {
			if errors.Is(err, postgres.ErrUserSegmentNotFound) {
				continue
			}
			return 0, err
		}
		log.Infof("Deleted user_id: %d from segment with id: %d and slug: %s", userId, segmentId, segmentSlugToDelete)
	}

	return userId, nil
}
