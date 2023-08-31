package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"segmenatationService/internal/api/resources"
	"segmenatationService/internal/models"
)

// @Summary Create segment
// @Description Create segment with given slug
// @Accept  json
// @Produce  json
// @Param data body models.Segment true "Segment"
// @Success 200 {object} []resources.SegmentResource
// @Failure 400 {object} resources.ErrorResponse
// @Router /segments [post]
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

// @Summary Delete segment by slug
// @Description Delete segment by given slug
// @Accept  json
// @Produce  json
// @Param slug path string true "Slug"
// @Success 200 {object} []resources.SegmentResource
// @Failure 400 {object} resources.ErrorResponse
// @Router /segments/{slug} [delete]
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
