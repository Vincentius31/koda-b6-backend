package handlers

import (
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

// GetAll godoc
// @Summary Get all users
// @Description Retrieve a list of all registered users
// @Tags users
// @Produce json
// @Success 200 {object} models.WebResponse
// @Failure 500 {object} models.WebResponse
// @Router /users [get]
func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to fetch users",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Successfully retrieved all users",
		Data:    users,
	})
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get detailed information of a specific user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.WebResponse
// @Failure 404 {object} models.WebResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid user ID format",
			Data:    nil,
		})
		return
	}

	user, err := h.service.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "User found",
		Data:    user,
	})
}

// Create godoc
// @Summary Register/Create new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "User Data"
// @Success 201 {object} models.WebResponse
// @Failure 400 {object} models.WebResponse
// @Failure 409 {object} models.WebResponse
// @Router /users [post]
func (h *UserHandler) Create(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if err := h.service.Register(ctx.Request.Context(), req); err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Email is already registered!" {
			statusCode = http.StatusConflict
		} else {
			statusCode = http.StatusBadRequest
		}

		ctx.JSON(statusCode, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true,
		Message: "User created successfully",
		Data:    nil,
	})
}

// Update godoc
// @Summary Update user (PATCH)
// @Description Update user data partially. Use PATCH method.
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body models.UpdateUserRequest true "Updated User Data"
// @Success 200 {object} models.WebResponse
// @Failure 404 {object} models.WebResponse
// @Router /users/{id} [patch]
func (h *UserHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid User ID",
			Data:    nil,
		})
		return
	}

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if err := h.service.Update(ctx.Request.Context(), id, req); err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete user
// @Description Remove user from database
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.WebResponse
// @Failure 500 {object} models.WebResponse
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.service.Remove(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to delete user",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "User deleted successfully",
		Data:    nil,
	})
}

// UploadProfile godoc
// @Summary Upload profile picture
// @Description Upload and update profile image for a user
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "User ID"
// @Param profile_image formData file true "Image File"
// @Success 200 {object} models.WebResponse
// @Router /users/{id}/upload [post]
func (h *UserHandler) UploadProfile(ctx *gin.Context) {
	userIdFromToken, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}

	// Konversi float64 (bawaan JSON/JWT) ke int
	id := int(userIdFromToken.(float64))

	file, err := ctx.FormFile("profile_image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "No file uploaded",
			Data:    nil,
		})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Only JPG, PNG, and JPEG are allowed",
			Data:    nil,
		})
		return
	}

	if file.Size > 1*1024*1024 {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "File size too large (Max 1MB)",
			Data:    nil,
		})
		return
	}

	filename := fmt.Sprintf("%d_%s", id, file.Filename)
	dst := "uploads/users/" + filename

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to save file. Make sure 'uploads/users/' folder exists!",
			Data:    nil,
		})
		return
	}

	err = h.service.UpdateProfileImage(ctx.Request.Context(), id, filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to update profile picture in database",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Profile picture uploaded successfully",
		Data:    map[string]string{"url": "/uploads/users/" + filename},
	})
}

// GetProfile godoc
// @Summary Get current user profile
// @Tags profile
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /profile [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userIdData, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{Success: false, Message: "Unauthorized", Data: nil})
		return
	}

	id := int(userIdData.(float64))

	user, err := h.service.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{Success: false, Message: "User not found", Data: nil})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Profile found",
		Data:    user,
	})
}

// UpdateProfile godoc
// @Summary Update current user profile
// @Tags profile
// @Accept json
// @Produce json
// @Param request body models.UpdateUserRequest true "Updated User Data"
// @Success 200 {object} models.WebResponse
// @Router /profile [patch]
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userIdData, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{Success: false, Message: "Unauthorized", Data: nil})
		return
	}

	id := int(userIdData.(float64))

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: err.Error(), Data: nil})
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
		Message: "Profile updated successfully",
		Data:    nil,
	})
}
