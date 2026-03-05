package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct{
	service *service.RoleService
}

func NewRoleHandler(s *service.RoleService) *RoleHandler{
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