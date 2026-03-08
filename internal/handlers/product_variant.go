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

func (h *ProductVariantHandler) Create(ctx *gin.Context) {
	var req models.CreateProductVariantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: err.Error(),
		})
		return
	}
	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true, 
		Message: "Variant created",
	})
}

func (h *ProductVariantHandler) GetAll(ctx *gin.Context) {
	data, _ := h.service.GetAll(ctx.Request.Context())
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Data: data,
	})
}

func (h *ProductVariantHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false, 
			Message: "Variant not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Data: data,
	})
}

func (h *ProductVariantHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var req models.UpdateProductVariantRequest
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
		Message: "Variant updated",
	})
}

func (h *ProductVariantHandler) Delete(ctx *gin.Context) {
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
		Message: "Variant deleted",
	})
}
