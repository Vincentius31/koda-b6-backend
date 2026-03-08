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

func (h *ProductImageHandler) Create(ctx *gin.Context) {
	productID, _ := strconv.Atoi(ctx.PostForm("product_id"))

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: "Image is required",
		})
		return
	}

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	savePath := filepath.Join("uploads/products", filename)
	
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: "Failed to save image",
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
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true, 
		Message: "Image uploaded successfully",
	})
}

func (h *ProductImageHandler) GetAll(ctx *gin.Context) {
	data, _ := h.service.GetAll(ctx.Request.Context())
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Data: data,
	})
}

func (h *ProductImageHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false, 
			Message: "Image not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Data: data,
	})
}

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
		ctx.SaveUploadedFile(file, savePath)
		req.Path = &savePath
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
		Message: "Image updated successfully",
	})
}

func (h *ProductImageHandler) Delete(ctx *gin.Context) {
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
		Message: "Image deleted",
	})
}