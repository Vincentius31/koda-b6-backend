package handlers

import (
	"github.com/gin-gonic/gin"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
)

type ProductPageHandler struct {
	service *service.ProductPageService
}

func NewProductPageHandler(s *service.ProductPageService) *ProductPageHandler {
	return &ProductPageHandler{service: s}
}

func (h *ProductPageHandler) GetCatalog(ctx *gin.Context) {
	params := map[string]string{
		"page":      ctx.DefaultQuery("page", "1"),
		"search":    ctx.Query("search"),
		"category":  ctx.Query("category"),
		"min_price": ctx.Query("min_price"),
		"max_price": ctx.Query("max_price"),
	}

	result, err := h.service.GetCatalogOnly(ctx.Request.Context(), params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to Fetch Products!",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Fetch Products Successfully!",
		Data:    result,
	})
}

func (h *ProductPageHandler) GetPromos(ctx *gin.Context) {
	result, err := h.service.GetAllPromos(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to Fetch Promos!",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Fetch Promos Successfully!",
		Data:    result,
	})
}
