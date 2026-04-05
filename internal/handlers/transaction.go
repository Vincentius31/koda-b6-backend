package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func getTransIDFromContext(ctx *gin.Context) (int, bool) {
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

// Checkout godoc
// @Summary Process checkout
// @Description Create transaction, add transaction products, and clear cart
// @Tags checkout
// @Accept json
// @Produce json
// @Param request body models.CheckoutRequest true "Checkout Data"
// @Success 201 {object} models.WebResponse
// @Failure 400 {object} models.WebResponse
// @Router /checkout [post]
func (h *TransactionHandler) Checkout(ctx *gin.Context) {
	userID, ok := getTransIDFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: "Unauthorized. Please login.",
		})
		return
	}

	var req models.CheckoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	result, err := h.service.Checkout(ctx.Request.Context(), userID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true,
		Message: "Checkout successful",
		Data:    result,
	})
}

// Create godoc
// @Summary Create a new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body models.CreateTransactionRequest true "Transaction Request"
// @Success 201 {object} models.WebResponse
// @Failure 400 {object} models.WebResponse
// @Router /transactions [post]
func (h *TransactionHandler) Create(ctx *gin.Context) {
	var req models.CreateTransactionRequest
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
		Message: "Transaction created successfully",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all transactions
// @Tags transactions
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /transactions [get]
func (h *TransactionHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to fetch transactions",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Successfully retrieved all transactions",
		Data:    data,
	})
}

// GetByID godoc
// @Summary Get transaction by ID
// @Tags transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.WebResponse
// @Failure 404 {object} models.WebResponse
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid ID format",
		})
		return
	}

	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Transaction not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Transaction found",
		Data:    data,
	})
}

// Update godoc
// @Summary Update transaction (PATCH)
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param request body models.UpdateTransactionRequest true "Update Request"
// @Success 200 {object} models.WebResponse
// @Router /transactions/{id} [patch]
func (h *TransactionHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateTransactionRequest
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
		Message: "Transaction updated successfully",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete transaction
// @Tags transactions
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.WebResponse
// @Router /transactions/{id} [delete]
func (h *TransactionHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to delete transaction",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Transaction deleted successfully",
		Data:    nil,
	})
}

// GetUserTransactions godoc
// @Summary Get all transactions for logged-in user
// @Tags transactions
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /transactions [get]
func (h *TransactionHandler) GetUserTransactions(ctx *gin.Context) {
	userID, ok := getTransIDFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: "Unauthorized. Please login.",
		})
		return
	}

	data, err := h.service.GetByUserID(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to fetch transactions",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Successfully retrieved transactions",
		Data:    data,
	})
}

// GetDetailByID godoc
// @Summary Get transaction detail by ID for logged-in user
// @Tags transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.WebResponse
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetDetailByID(ctx *gin.Context) {
	userID, ok := getTransIDFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: "Unauthorized. Please login.",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid ID format",
		})
		return
	}

	data, err := h.service.GetDetailByID(ctx.Request.Context(), id, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Transaction not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Transaction found",
		Data:    data,
	})
}
