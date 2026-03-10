package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductVariantHandler struct {
	service *service.ProductVariantService
}

func NewProductVariantHandler(s *service.ProductVariantService) *ProductVariantHandler {
	return &ProductVariantHandler{service: s}
}

// Create godoc
// @Summary Create product variant
// @Tags product-variants
// @Accept json
// @Produce json
// @Param request body models.CreateProductVariantRequest true "Variant Data"
// @Success 201 {object} models.WebResponse
// @Failure 400 {object} models.WebResponse
// @Router /productvariant [post]
func (h *ProductVariantHandler) Create(ctx *gin.Context) {
	var req models.CreateProductVariantRequest
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
		Message: "Variant created successfully",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all product variants
// @Tags product-variants
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /productvariant [get]
func (h *ProductVariantHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to fetch variants",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Variants retrieved successfully",
		Data:    data,
	})
}

// GetByID godoc
// @Summary Get variant by ID
// @Tags product-variants
// @Produce json
// @Param id path int true "Variant ID"
// @Success 200 {object} models.WebResponse
// @Router /productvariant/{id} [get]
func (h *ProductVariantHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Variant not found",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Variant found",
		Data:    data,
	})
}

// Update godoc
// @Summary Update variant (PATCH)
// @Tags product-variants
// @Accept json
// @Produce json
// @Param id path int true "Variant ID"
// @Param request body models.UpdateProductVariantRequest true "Update Data"
// @Success 200 {object} models.WebResponse
// @Router /productvariant/{id} [patch]
func (h *ProductVariantHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateProductVariantRequest
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
		Message: "Variant updated successfully",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete variant
// @Tags product-variants
// @Param id path int true "Variant ID"
// @Success 200 {object} models.WebResponse
// @Router /productvariant/{id} [delete]
func (h *ProductVariantHandler) Delete(ctx *gin.Context) {
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
		Message: "Variant deleted successfully",
		Data:    nil,
	})
}
