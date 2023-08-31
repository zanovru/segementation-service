package models

type SegmentProb struct {
	Id          int     `json:"id" db:"id"`
	Slug        string  `json:"slug" db:"slug" binding:"required"`
	Probability float64 `json:"probability" binding:"required,gte=0,lte=1"`
}

func (sp *SegmentProb) ToSegmentModel() *Segment {
	return &Segment{
		Slug: sp.Slug,
	}
}
