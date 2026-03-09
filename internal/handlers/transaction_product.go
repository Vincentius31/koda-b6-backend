package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionProductHandler struct {
	service *service.TransactionProductService
}

func NewTransactionProductHandler(s *service.TransactionProductService) *TransactionProductHandler {
	return &TransactionProductHandler{service: s}
}

// Create godoc
// @Summary Add item to transaction
// @Tags transaction-products
// @Accept json
// @Produce json
// @Param request body models.CreateTransactionProductRequest true "Item Data"
// @Success 201 {object} models.WebResponse
// @Failure 400 {object} models.WebResponse
// @Router /transactionproduct [post]
func (h *TransactionProductHandler) Create(ctx *gin.Context) {
	var req models.CreateTransactionProductRequest
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
		Message: "Item added to transaction",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all transaction items
// @Tags transaction-products
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /transactionproduct [get]
func (h *TransactionProductHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to fetch items",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Successfully retrieved items",
		Data:    data,
	})
}

// GetByID godoc
// @Summary Get transaction item by ID
// @Tags transaction-products
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.WebResponse
// @Failure 404 {object} models.WebResponse
// @Router /transactionproduct/{id} [get]
func (h *TransactionProductHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid ID format",
			Data:    nil,
		})
		return
	}

	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Item not found",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Item found",
		Data:    data,
	})
}

// Update godoc
// @Summary Update transaction item (PATCH)
// @Tags transaction-products
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param request body models.UpdateTransactionProductRequest true "Updated Data"
// @Success 200 {object} models.WebResponse
// @Router /transactionproduct/{id} [patch]
func (h *TransactionProductHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateTransactionProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Transaction item updated",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete transaction item
// @Tags transaction-products
// @Param id path int true "Item ID"
// @Success 200 {object} models.WebResponse
// @Router /transactionproduct/{id} [delete]
func (h *TransactionProductHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Item removed from transaction",
		Data:    nil,
	})
}
