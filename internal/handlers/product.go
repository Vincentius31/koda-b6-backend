package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

// Create godoc
// @Summary Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param request body models.CreateProductRequest true "Product Data"
// @Success 201 {object} models.WebResponse
// @Router /product [post]
func (h *ProductHandler) Create(ctx *gin.Context) {
	var req models.CreateProductRequest
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
		Message: "Product created successfully",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all products
// @Tags products
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /product [get]
func (h *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to retrieve products",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Products retrieved",
		Data:    products,
	})
}

// GetByID godoc
// @Summary Get product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.WebResponse
// @Router /product/{id} [get]
func (h *ProductHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: "Invalid ID", 
			Data: nil,
		})
		return
	}

	product, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false, 
			Message: "Product not found", 
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Product found", 
		Data: product,
	})
}

// Update godoc
// @Summary Update product (PATCH)
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body models.UpdateProductRequest true "Update Data"
// @Success 200 {object} models.WebResponse
// @Router /product/{id} [patch]
func (h *ProductHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateProductRequest
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
		Message: "Product updated successfully", 
		Data: nil,
	})
}

// Delete godoc
// @Summary Delete product
// @Tags products
// @Param id path int true "Product ID"
// @Success 200 {object} models.WebResponse
// @Router /product/{id} [delete]
func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: "Invalid ID", 
			Data: nil,
		})
		return
	}

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
		Message: "Product deleted successfully", 
		Data: nil,
	})
}
