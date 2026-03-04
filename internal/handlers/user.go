package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	user, err := h.service.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Register(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Registration failed",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated",
	})
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Remove(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}