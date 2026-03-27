package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DetailProductHandler struct {
	service *service.DetailProductService
}

func NewDetailProductHandler(s *service.DetailProductService) *DetailProductHandler {
	return &DetailProductHandler{service: s}
}

func (h *DetailProductHandler) GetDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid Product ID!",
		})
		return
	}

	result, err := h.service.GetDetailByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Product Not Found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Product Detail Fetched Successfully!",
		Data:    result,
	})
}