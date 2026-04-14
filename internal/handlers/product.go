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

func (h *ProductHandler) parseProductForm(ctx *gin.Context) (models.AdminProductPayload, error) {
	var payload models.AdminProductPayload

	payload.NameProduct = ctx.PostForm("nameProduct")
	payload.Description = ctx.PostForm("description")
	payload.Category = ctx.PostForm("category")
	payload.PromoType = ctx.PostForm("promoType")

	payload.PriceProduct, _ = strconv.Atoi(ctx.PostForm("priceProduct"))
	payload.PriceDiscount, _ = strconv.Atoi(ctx.PostForm("priceDiscount"))
	payload.Stock, _ = strconv.Atoi(ctx.PostForm("stock"))

	payload.Size = ctx.PostFormArray("size")
	if len(payload.Size) == 0 && ctx.PostForm("size") != "" {
		payload.Size = strings.Split(ctx.PostForm("size"), ",")
	}

	payload.Temp = ctx.PostFormArray("temp")
	if len(payload.Temp) == 0 && ctx.PostForm("temp") != "" {
		payload.Temp = strings.Split(ctx.PostForm("temp"), ",")
	}

	payload.Method = ctx.PostFormArray("method")
	if len(payload.Method) == 0 && ctx.PostForm("method") != "" {
		payload.Method = strings.Split(ctx.PostForm("method"), ",")
	}

	return payload, nil
}

func (h *ProductHandler) handleFileUpload(ctx *gin.Context) ([]string, error) {
	var savedPaths []string

	form, err := ctx.MultipartForm()
	if err != nil {
		return savedPaths, nil
	}

	files := form.File["images"]
	if len(files) == 0 {
		return savedPaths, nil
	}

	os.MkdirAll("uploads", os.ModePerm)

	for _, file := range files {
		extension := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
		filePath := filepath.Join("uploads", newFileName)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			return nil, err
		}

		savedPaths = append(savedPaths, "http://localhost:8080/"+filePath)
	}

	return savedPaths, nil
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	payload, err := h.parseProductForm(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid form data"})
		return
	}

	imagePaths, err := h.handleFileUpload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: "Failed to save images"})
		return
	}
	payload.ImageProduct = imagePaths

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

	payload, err := h.parseProductForm(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{Success: false, Message: "Invalid form data"})
		return
	}

	imagePaths, err := h.handleFileUpload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: "Failed to upload new images"})
		return
	}
	
	if len(imagePaths) > 0 {
		payload.ImageProduct = imagePaths
	}

	if err := h.service.Update(ctx.Request.Context(), id, payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{Success: false, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{Success: true, Message: "Product updated successfully"})
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