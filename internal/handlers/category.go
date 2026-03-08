package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var req models.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: err.Error(),
		})
		return
	}
	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true, 
		Message: "Category created",
	})
}

func (h *CategoryHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Data: data,
	})
}

func (h *CategoryHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false, 
			Message: "Category not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Data: data,
	})
}

func (h *CategoryHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: err.Error(),
		})
		return
	}
	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Category updated",
	})
}

func (h *CategoryHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Category deleted",
	})
}