package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service *service.ReviewService
}

func NewReviewHandler(s *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: s}
}

// Create godoc
// @Summary Submit a new review
// @Tags reviews
// @Accept json
// @Produce json
// @Param request body models.CreateReviewRequest true "Review Data"
// @Success 201 {object} models.WebResponse
// @Failure 400 {object} models.WebResponse
// @Router /review [post]
func (h *ReviewHandler) Create(ctx *gin.Context) {
	var req models.CreateReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true,
		Message: "Review submitted successfully",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all reviews
// @Tags reviews
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /review [get]
func (h *ReviewHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to fetch reviews",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Successfully retrieved all reviews",
		Data:    data,
	})
}

// GetByID godoc
// @Summary Get review by ID
// @Tags reviews
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} models.WebResponse
// @Failure 404 {object} models.WebResponse
// @Router /review/{id} [get]
func (h *ReviewHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Review not found",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Review found",
		Data:    data,
	})
}

// Update godoc
// @Summary Update review (PATCH)
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param request body models.UpdateReviewRequest true "Updated Review Data"
// @Success 200 {object} models.WebResponse
// @Router /review/{id} [patch]
func (h *ReviewHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Review updated successfully",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete review
// @Tags reviews
// @Param id path int true "Review ID"
// @Success 200 {object} models.WebResponse
// @Router /review/{id} [delete]
func (h *ReviewHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to delete review",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Review deleted successfully",
		Data:    nil,
	})
}
