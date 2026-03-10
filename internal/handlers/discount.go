package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	service *service.DiscountService
}

func NewDiscountHandler(s *service.DiscountService) *DiscountHandler {
	return &DiscountHandler{service: s}
}

// Create godoc
// @Summary Create a new discount
// @Tags discounts
// @Accept json
// @Produce json
// @Param request body models.CreateDiscountRequest true "Discount Data"
// @Success 201 {object} models.WebResponse
// @Router /discount [post]
func (h *DiscountHandler) Create(ctx *gin.Context) {
	var req models.CreateDiscountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true, 
		Message: "Discount created successfully", 
		Data: nil,
	})
}

// GetAll godoc
// @Summary Get all discounts
// @Tags discounts
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /discount [get]
func (h *DiscountHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Discounts retrieved", 
		Data: data,
	})
}

// GetByID godoc
// @Summary Get discount by ID
// @Tags discounts
// @Produce json
// @Param id path int true "Discount ID"
// @Success 200 {object} models.WebResponse
// @Router /discount/{id} [get]
func (h *DiscountHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false, 
			Message: "Discount not found", 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Discount found", 
		Data: data,
	})
}

// Update godoc
// @Summary Update discount (PATCH)
// @Tags discounts
// @Accept json
// @Produce json
// @Param id path int true "Discount ID"
// @Param request body models.UpdateDiscountRequest true "Update Data"
// @Success 200 {object} models.WebResponse
// @Router /discount/{id} [patch]
func (h *DiscountHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateDiscountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}

	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Discount updated successfully", 
		Data: nil,
	})
}

// Delete godoc
// @Summary Delete discount
// @Tags discounts
// @Param id path int true "Discount ID"
// @Success 200 {object} models.WebResponse
// @Router /discount/{id} [delete]
func (h *DiscountHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Discount deleted successfully", 
		Data: nil,
	})
}
