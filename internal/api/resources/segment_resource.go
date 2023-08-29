package resources

type SegmentResource struct {
	Id int `json:"id"`
}

func NewSegmentResource(id int) *SegmentResource {
	return &SegmentResource{Id: id}
}
