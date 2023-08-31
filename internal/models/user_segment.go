package models

type UserSegment struct {
	UserId    int `json:"userId" db:"user_id"`
	SegmentId int `json:"segmentId" db:"segment_id"`
}
