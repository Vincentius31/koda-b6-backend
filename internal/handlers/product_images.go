package handlers

import (
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductImageHandler struct {
	service *service.ProductImageService
}

func NewProductImageHandler(s *service.ProductImageService) *ProductImageHandler {
	return &ProductImageHandler{service: s}
}

// Create godoc
// @Summary Upload product image
// @Tags product-images
// @Accept multipart/form-data
// @Produce json
// @Param product_id formData int true "Product ID"
// @Param image formData file true "Image File"
// @Success 201 {object} models.WebResponse
// @Router /productimage [post]
func (h *ProductImageHandler) Create(ctx *gin.Context) {
	productID, _ := strconv.Atoi(ctx.PostForm("product_id"))

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Image is required",
			Data:    nil,
		})
		return
	}

	// Gunakan timestamp untuk menghindari nama file duplikat
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	savePath := filepath.Join("uploads/products", filename)

	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: "Failed to save image",
			Data:    nil,
		})
		return
	}

	img := models.ProductImage{
		ProductID: productID,
		Path:      savePath,
	}

	if err := h.service.Create(ctx.Request.Context(), img); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true,
		Message: "Image uploaded successfully",
		Data:    nil,
	})
}

// GetAll godoc
// @Summary Get all product images
// @Tags product-images
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /productimage [get]
func (h *ProductImageHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
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
		Message: "Images retrieved successfully",
		Data:    data,
	})
}

// GetByID godoc
// @Summary Get product image by ID
// @Tags product-images
// @Produce json
// @Param id path int true "Image ID"
// @Success 200 {object} models.WebResponse
// @Failure 404 {object} models.WebResponse
// @Router /productimage/{id} [get]
func (h *ProductImageHandler) GetByID(ctx *gin.Context) {
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
            Message: "Image not found",
            Data:    nil,
        })
        return
    }

    ctx.JSON(http.StatusOK, models.WebResponse{
        Success: true,
        Message: "Image found",
        Data:    data,
    })
}

// Update godoc
// @Summary Update product image (PATCH)
// @Tags product-images
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Image ID"
// @Param product_id formData int false "Product ID"
// @Param image formData file false "Image File"
// @Success 200 {object} models.WebResponse
// @Router /productimage/{id} [patch]
func (h *ProductImageHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	pIDStr := ctx.PostForm("product_id")
	var req models.UpdateProductImageRequest

	if pIDStr != "" {
		pID, _ := strconv.Atoi(pIDStr)
		req.ProductID = &pID
	}

	file, err := ctx.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
		savePath := filepath.Join("uploads/products", filename)
		if err := ctx.SaveUploadedFile(file, savePath); err == nil {
			req.Path = &savePath
		}
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
		Message: "Image updated successfully",
		Data:    nil,
	})
}

// Delete godoc
// @Summary Delete product image
// @Tags product-images
// @Param id path int true "Image ID"
// @Success 200 {object} models.WebResponse
// @Router /productimage/{id} [delete]
func (h *ProductImageHandler) Delete(ctx *gin.Context) {
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
		Message: "Image deleted successfully",
		Data:    nil,
	})
}