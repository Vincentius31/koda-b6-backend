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

func (h *DiscountHandler) Create(ctx *gin.Context) {
	var req models.CreateDiscountRequest
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
		Message: "Discount created",
	})
}

func (h *DiscountHandler) GetAll(ctx *gin.Context) {
	data, _ := h.service.GetAll(ctx.Request.Context())
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Data:    data,
	})
}

func (h *DiscountHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Discount not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Data:    data,
	})
}

func (h *DiscountHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var req models.UpdateDiscountRequest
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
		Message: "Discount updated",
	})
}

func (h *DiscountHandler) Delete(ctx *gin.Context) {
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
		Message: "Discount deleted",
	})
}
