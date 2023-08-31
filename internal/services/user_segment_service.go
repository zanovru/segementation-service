package services

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"segmenatationService/internal/contracts"
	"segmenatationService/internal/models"
	"time"
)

type UserSegmentService struct {
	userRepo        contracts.UserRepositoryContract
	userSegmentRepo contracts.UserSegmentRepositoryContract
}

func NewUserSegmentService(userRepo contracts.UserRepositoryContract,
	userSegmentRepo contracts.UserSegmentRepositoryContract) *UserSegmentService {
	return &UserSegmentService{userRepo: userRepo, userSegmentRepo: userSegmentRepo}
}

func (u UserSegmentService) CreateUserSegmentWithProb(segmentId int, prob float64) error {
	users, err := u.userRepo.GetUsers()
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	for _, user := range users {
		random := rand.Float64()
		if random < prob {
			log.Info("generated: ", random, prob)
			userSegment := &models.UserSegment{
				UserId:    user.Id,
				SegmentId: segmentId,
			}
			userId, segId, err := u.userSegmentRepo.CreateUserSegment(userSegment)
			if err != nil {
				return err
			}
			log.Infof("Added user_id: %d to segment with id: %d ", userId, segId)
		}
	}
	return nil
}
