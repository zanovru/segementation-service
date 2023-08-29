package contracts

import "segmenatationService/internal/models"

type SegmentRepositoryContract interface {
	CreateSegment(*models.Segment) (int, error)
	DeleteSegmentBySlug(slug string) (int, error)
}
