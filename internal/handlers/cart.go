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

// Create godoc
// @Summary Add item to cart
// @Tags carts
// @Accept json
// @Produce json
// @Param request body models.CreateCartRequest true "Cart Item Data"
// @Success 201 {object} models.WebResponse
// @Router /cart [post]
func (h *CartHandler) Create(ctx *gin.Context) {
	var req models.CreateCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{Success: true, Message: "Item added to cart", Data: nil})
}

// GetAll godoc
// @Summary Get all cart items
// @Tags carts
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /cart [get]
func (h *CartHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Cart retrieved successfully", Data: data})
}

// GetByID godoc
// @Summary Get cart item by ID
// @Tags carts
// @Produce json
// @Param id path int true "Cart ID"
// @Success 200 {object} models.WebResponse
// @Router /cart/{id} [get]
func (h *CartHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{Success: false, Message: "Cart item not found", Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Cart item found", Data: data})
}

// Update godoc
// @Summary Update cart item (PATCH)
// @Tags carts
// @Accept json
// @Produce json
// @Param id path int true "Cart ID"
// @Param request body models.UpdateCartRequest true "Update Data"
// @Success 200 {object} models.WebResponse
// @Router /cart/{id} [patch]
func (h *CartHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Cart updated successfully", Data: nil})
}

// Delete godoc
// @Summary Remove item from cart
// @Tags carts
// @Param id path int true "Cart ID"
// @Success 200 {object} models.WebResponse
// @Router /cart/{id} [delete]
func (h *CartHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Item removed from cart", Data: nil})
}
