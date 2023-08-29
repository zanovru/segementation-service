package models

type User struct {
	Name             string    `json:"name"`
	SegmentsToAdd    []Segment `json:"segmentsToAdd"`
	SegmentsToDelete []Segment `json:"segmentsToDelete"`
}
