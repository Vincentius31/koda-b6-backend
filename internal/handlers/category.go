package handlers

import (
	"github.com/gin-gonic/gin"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

// Create godoc
// @Summary Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param request body models.CreateCategoryRequest true "Category Data"
// @Success 201 {object} models.WebResponse
// @Router /category [post]
func (h *CategoryHandler) Create(ctx *gin.Context) {
	var req models.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{Success: true, Message: "Category created successfully", Data: nil})
}

// GetAll godoc
// @Summary Get all categories
// @Tags categories
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /category [get]
func (h *CategoryHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Categories retrieved", Data: data})
}

// GetByID godoc
// @Summary Get category by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.WebResponse
// @Router /category/{id} [get]
func (h *CategoryHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{Success: false, Message: "Category not found", Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Category found", Data: data})
}

// Update godoc
// @Summary Update category (PATCH)
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param request body models.UpdateCategoryRequest true "Category Data"
// @Success 200 {object} models.WebResponse
// @Router /category/{id} [patch]
func (h *CategoryHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Category updated successfully", Data: nil})
}

// Delete godoc
// @Summary Delete category
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} models.WebResponse
// @Router /category/{id} [delete]
func (h *CategoryHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Category deleted successfully", Data: nil})
}
