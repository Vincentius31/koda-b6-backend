package handlers

import (
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetPromos(ctx *gin.Context) {
	promos, err := h.service.GetPromos(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: "Failed to retrieve promos"})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Promos retrieved", Data: promos})
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var payload models.AdminProductPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid JSON data"})
		return
	}

	if err := h.service.Create(ctx.Request.Context(), payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{Success: true, Message: "Product created successfully"})
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: "Failed to retrieve products"})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Products retrieved", Data: products})
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid ID"})
		return
	}

	product, err := h.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{Success: false, Message: "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Product found", Data: product})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid ID"})
		return
	}

	var payload models.AdminProductPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid JSON data"})
		return
	}

	if err := h.service.Update(ctx.Request.Context(), id, payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Product updated successfully"})
}

func (h *ProductHandler) UploadImages(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid ID"})
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Failed to parse form"})
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "No images uploaded"})
		return
	}

	os.MkdirAll("uploads/products", os.ModePerm)

	var savedPaths []string
	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Only JPG and PNG allowed"})
			return
		}
		newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join("uploads/products", newFileName)
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: "Failed to save image"})
			return
		}
		savedPaths = append(savedPaths, "/"+filePath)
	}

	if err := h.service.UpdateImages(ctx.Request.Context(), id, savedPaths); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error()})
		return
	}

	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Images uploaded successfully", Data: savedPaths})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid ID"})
		return
	}

	if err := h.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Product deleted successfully"})
}
