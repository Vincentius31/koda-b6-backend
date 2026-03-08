package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	service *service.RoleService
}

func NewRoleHandler(s *service.RoleService) *RoleHandler {
	return &RoleHandler{
		service: s,
	}
}

func (h *RoleHandler) Create(ctx *gin.Context) {
	var req models.CreateRoleRequest
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
		Message: "Role created successfully",
	})
}

func (h *RoleHandler) GetAll(ctx *gin.Context) {
	roles, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Roles retrieved", Data: roles,
	})
}

func (h *RoleHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	role, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Role found",
		Data:    role,
	})
}

func (h *RoleHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateRoleRequest
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
		Message: "Role updated successfully",
	})
}

func (h *RoleHandler) Delete(ctx *gin.Context) {
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
		Message: "Role deleted successfully",
	})
}
