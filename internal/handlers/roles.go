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
	return &RoleHandler{service: s}
}

// Create godoc
// @Summary Create a new role
// @Tags roles
// @Accept json
// @Produce json
// @Param request body models.CreateRoleRequest true "Role Name"
// @Success 201 {object} models.WebResponse
// @Failure 400 {object} models.WebResponse
// @Router /roles [post]
func (h *RoleHandler) Create(ctx *gin.Context) {
	var req models.CreateRoleRequest
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
		Message: "Role created successfully",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all roles
// @Tags roles
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /roles [get]
func (h *RoleHandler) GetAll(ctx *gin.Context) {
	roles, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Roles retrieved",
		Data:    roles,
	})
}

// GetByID godoc
// @Summary Get role by ID
// @Tags roles
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} models.WebResponse
// @Failure 404 {object} models.WebResponse
// @Router /roles/{id} [get]
func (h *RoleHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid ID format",
			Data:    nil,
		})
		return
	}

	role, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Role found",
		Data:    role,
	})
}

// Update godoc
// @Summary Update role (PATCH)
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param request body models.UpdateRoleRequest true "Updated Role Name"
// @Success 200 {object} models.WebResponse
// @Router /roles/{id} [patch]
func (h *RoleHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req models.UpdateRoleRequest
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
		Message: "Role updated successfully",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete role
// @Tags roles
// @Param id path int true "Role ID"
// @Success 200 {object} models.WebResponse
// @Router /roles/{id} [delete]
func (h *RoleHandler) Delete(ctx *gin.Context) {
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
		Message: "Role deleted successfully",
		Data:    nil,
	})
}
