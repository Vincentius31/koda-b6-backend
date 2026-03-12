package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LandingHandler struct {
	landingService *service.LandingService
}

func NewLandingHandler(ls *service.LandingService) *LandingHandler {
	return &LandingHandler{
		landingService: ls,
	}
}

func (h *LandingHandler) GetRecommendedProducts(ctx *gin.Context) {
	products, err := h.landingService.GetRecommendedProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to load recommended products",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Success to load recommended products",
		Data:    products,
	})
}

func (h *LandingHandler) GetLatestReviews(ctx *gin.Context) {
	reviews, err := h.landingService.GetLatestReviews(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to retrieve review data",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Successfully retrieved the latest review",
		Data:    reviews,
	})
}
