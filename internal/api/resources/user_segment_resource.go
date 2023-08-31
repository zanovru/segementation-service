package resources

import "segmenatationService/internal/models"

type UserSegmentResource struct {
	Id       int              `json:"id"`
	Segments []models.Segment `json:"segments"`
}

func NewUserSegmentResource(id int, segments []models.Segment) *UserSegmentResource {
	return &UserSegmentResource{Id: id, Segments: segments}
}
