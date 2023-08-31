package models

type Segment struct {
	Id   int    `json:"id" db:"id"`
	Slug string `json:"slug" db:"slug" binding:"required"`
}
