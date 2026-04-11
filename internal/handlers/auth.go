package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService   *service.UserService
	forgotService *service.ForgotPasswordService
}

func NewAuthHandler(us *service.UserService, fs *service.ForgotPasswordService) *AuthHandler {
	return &AuthHandler{
		userService:   us,
		forgotService: fs,
	}
}

// Login godoc
// @Summary User Login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login Credentials"
// @Success 200 {object} models.WebResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Email and password are required",
			Data:    nil,
		})
		return
	}

	token, roleID, err := h.userService.Login(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Login successful",
		Data: models.LoginResponse{
			Token:  token,
			RoleID: roleID,
		},
	})
}

// Register godoc
// @Summary Register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "User Data"
// @Success 201 {object} models.WebResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if err := h.userService.Register(ctx.Request.Context(), req); err != nil {
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

// RequestForgotPassword godoc
// @Summary Request OTP for password reset
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.ForgotPasswordRequest true "Email Data"
// @Success 200 {object} models.WebResponse
// @Router /auth/forgot-password [post]
func (h *AuthHandler) RequestForgotPassword(ctx *gin.Context) {
	var req models.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid input: Email is required",
			Data:    nil,
		})
		return
	}

	err := h.forgotService.RequestForgotPassword(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Email is not registered in our system",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "OTP code has been sent.",
		Data:    nil,
	})
}

// ResetPassword godoc
// @Summary Reset password using OTP
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.ResetPasswordRequest true "Reset Data"
// @Success 200 {object} models.WebResponse
// @Router /auth/reset-password [patch]
func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req models.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Invalid input data",
			Data:    nil,
		})
		return
	}

	err := h.forgotService.ResetPassword(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.WebResponse{
			Success: false,
			Message: "Invalid OTP code or email",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Password has been updated successfully",
		Data:    nil,
	})
}
