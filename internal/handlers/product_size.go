package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductSizeHandler struct {
	service *service.ProductSizeService
}

func NewProductSizeHandler(s *service.ProductSizeService) *ProductSizeHandler {
	return &ProductSizeHandler{service: s}
}

// Create godoc
// @Summary Add new product size
// @Tags product-sizes
// @Accept json
// @Produce json
// @Param request body models.CreateProductSizeRequest true "Size Data"
// @Success 201 {object} models.WebResponse
// @Router /productsize [post]
func (h *ProductSizeHandler) Create(ctx *gin.Context) {
	var req models.CreateProductSizeRequest
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
		Message: "Size created successfully",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all product sizes
// @Tags product-sizes
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /productsize [get]
func (h *ProductSizeHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to retrieve sizes",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Sizes retrieved successfully",
		Data:    data,
	})
}

// GetByID godoc
// @Summary Get product size by ID
// @Tags product-sizes
// @Produce json
// @Param id path int true "Size ID"
// @Success 200 {object} models.WebResponse
// @Router /productsize/{id} [get]
func (h *ProductSizeHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Size not found",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Size found",
		Data:    data,
	})
}

// Update godoc
// @Summary Update product size (PATCH)
// @Tags product-sizes
// @Accept json
// @Produce json
// @Param id path int true "Size ID"
// @Param request body models.UpdateProductSizeRequest true "Updated Data"
// @Success 200 {object} models.WebResponse
// @Router /productsize/{id} [patch]
func (h *ProductSizeHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateProductSizeRequest
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
		Message: "Size updated successfully",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete product size
// @Tags product-sizes
// @Param id path int true "Size ID"
// @Success 200 {object} models.WebResponse
// @Router /productsize/{id} [delete]
func (h *ProductSizeHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Size deleted successfully",
		Data:    nil,
	})
}
