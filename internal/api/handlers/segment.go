package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"segmenatationService/internal/api/resources"
	"segmenatationService/internal/models"
)

func (r *Routing) createSegment(c *gin.Context) {
	segment := &models.Segment{}
	if err := c.BindJSON(segment); err != nil {
		resources.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	segmentId, err := r.service.CreateSegment(segment)
	if err != nil {
		resources.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	segmentResource := resources.NewSegmentResource(segmentId)

	c.JSON(http.StatusCreated, segmentResource)
}

func (r *Routing) deleteSegmentBySlug(c *gin.Context) {
	slug := c.Params.ByName("slug")

	id, err := r.service.DeleteSegmentBySlug(slug)
	if err != nil {
		resources.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	segmentResource := resources.NewSegmentResource(id)

	c.JSON(http.StatusOK, segmentResource)
}
