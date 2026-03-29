package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(s *service.CartService) *CartHandler {
	return &CartHandler{service: s}
}

func getIDFromContext(ctx *gin.Context) (int, bool) {
	val, exists := ctx.Get("user_id")
	if !exists {
		val, exists = ctx.Get("userID")
	}

	if !exists {
		return 0, false
	}

	switch v := val.(type) {
	case int:
		return v, true
	case float64:
		return int(v), true
	case string:
		id, _ := strconv.Atoi(v)
		return id, true
	default:
		return 0, false
	}
}

func (h *CartHandler) Create(ctx *gin.Context) {
	userID, ok := getIDFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: "Unauthorized. Please login.",
		})
		return
	}

	var req models.CreateCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	req.UserID = userID

	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to add to cart: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true,
		Message: "Item added to cart",
	})
}

func (h *CartHandler) GetUserCart(ctx *gin.Context) {
	userID, ok := getIDFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: "Unauthorized access",
		})
		return
	}

	data, err := h.service.GetUserCart(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Cart retrieved successfully",
		Data:    data,
	})
}

func (h *CartHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid cart ID format",
		})
		return
	}

	var req models.UpdateCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.UpdateQty(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Cart updated",
	})
}

func (h *CartHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid cart ID format",
		})
		return
	}

	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Item removed from cart",
	})
}