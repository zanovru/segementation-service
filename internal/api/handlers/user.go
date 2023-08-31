package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"segmenatationService/internal/api/resources"
	"segmenatationService/internal/models"
	"strconv"
)

// @Summary Create user
// @Description Create user with segments that should be added to User and segments that should be deleted from User
// @Accept  json
// @Produce  json
// @Param data body models.User true "User"
// @Success 200 {object} []resources.UserResource
// @Failure 400 {object} resources.ErrorResponse
// @Router /users [post]
func (r *Routing) createUserWithSegments(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(user); err != nil {
		resources.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := r.service.CreateUserWithSegments(user)
	if err != nil {
		resources.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userResource := resources.NewUserResource(userId)
	c.JSON(http.StatusCreated, userResource)
}

// @Summary Get user segments
// @Description Get segments that were added to user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} []resources.UserSegmentResource
// @Failure 400 {object} resources.ErrorResponse
// @Router /users/{id}/segments [get]
func (r *Routing) getUserSegments(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		resources.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := r.service.GetUser(userId)
	if err != nil {
		resources.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	segments, err := r.service.GetSegmentsByUserId(user.Id)

	userSegmentResource := resources.NewUserSegmentResource(user.Id, segments)

	c.JSON(http.StatusOK, userSegmentResource)

}
