package models

type User struct {
	Id               int      `json:"id" db:"id"`
	Email            string   `json:"email" db:"email" binding:"required,email"`
	SegmentsToAdd    []string `json:"segmentsToAdd"`
	SegmentsToDelete []string `json:"segmentsToDelete"`
}
