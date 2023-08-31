package contracts

import "segmenatationService/internal/models"

type SegmentServiceContract interface {
	CreateSegment(*models.Segment) (int, error)
	DeleteSegmentBySlug(string) (int, error)
	GetSegmentsByUserId(id int) ([]models.Segment, error)
}

type UserServiceContract interface {
	CreateUserWithSegments(user *models.User) (int, error)
	GetUser(int) (*models.User, error)
}
