package contracts

import "segmenatationService/internal/models"

type SegmentRepositoryContract interface {
	CreateSegment(*models.Segment) (int, error)
	GetSegmentBySlug(string) (*models.Segment, error)
	DeleteSegmentBySlug(string) (int, error)
	GetSegmentsByUserId(id int) ([]models.Segment, error)
}

type UserRepositoryContract interface {
	CreateUser(*models.User) (int, error)
	GetUserByEmail(string) (*models.User, error)
	GetUser(int) (*models.User, error)
	GetUsers() ([]models.User, error)
}

type UserSegmentRepositoryContract interface {
	CreateUserSegment(*models.UserSegment) (int, int, error)
	DeleteUserSegment(*models.UserSegment) (int, int, error)
}
