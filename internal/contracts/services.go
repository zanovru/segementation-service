package contracts

import "segmenatationService/internal/models"

type SegmentServiceContract interface {
	CreateSegment(*models.Segment) (int, error)
	DeleteSegmentBySlug(string) (int, error)
}
