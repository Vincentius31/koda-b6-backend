package handlers

import (
	"github.com/gin-gonic/gin"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"strconv"
)

type ProductCategoryHandler struct {
	service *service.ProductCategoryService
}

func NewProductCategoryHandler(s *service.ProductCategoryService) *ProductCategoryHandler {
	return &ProductCategoryHandler{service: s}
}

// Create godoc
// @Summary Create product-category relation
// @Tags product-categories
// @Accept json
// @Produce json
// @Param request body models.CreateProductCategoryRequest true "Relation Data"
// @Success 201 {object} models.WebResponse
// @Router /productcategory [post]
func (h *ProductCategoryHandler) Create(ctx *gin.Context) {
	var req models.CreateProductCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	if err := h.service.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.WebResponse{
		Success: true, 
		Message: "Relation created successfully", 
		Data: nil,
	})
}

// GetAll godoc
// @Summary Get all product-category relations
// @Tags product-categories
// @Produce json
// @Success 200 {object} models.WebResponse
// @Router /productcategory [get]
func (h *ProductCategoryHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Relations retrieved", 
		Data: data,
	})
}

// GetByID godoc
// @Summary Get relation by product and category ID
// @Tags product-categories
// @Produce json
// @Param product_id query int true "Product ID"
// @Param category_id query int true "Category ID"
// @Success 200 {object} models.WebResponse
// @Router /productcategory/detail [get]
func (h *ProductCategoryHandler) GetByID(ctx *gin.Context) {
	pID, _ := strconv.Atoi(ctx.Query("product_id"))
	cID, _ := strconv.Atoi(ctx.Query("category_id"))

	data, err := h.service.GetByID(ctx.Request.Context(), pID, cID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false, 
			Message: "Relation not found", 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Relation found", 
		Data: data,
	})
}

// Update godoc
// @Summary Update relation (PATCH)
// @Tags product-categories
// @Accept json
// @Produce json
// @Param old_product_id query int true "Old Product ID"
// @Param old_category_id query int true "Old Category ID"
// @Param request body models.UpdateProductCategoryRequest true "New Relation Data"
// @Success 200 {object} models.WebResponse
// @Router /productcategory [patch]
func (h *ProductCategoryHandler) Update(ctx *gin.Context) {
	oldPID, errP := strconv.Atoi(ctx.Query("old_product_id"))
	oldCID, errC := strconv.Atoi(ctx.Query("old_category_id"))

	if errP != nil || errC != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: "Old Product ID and Category ID are required", 
			Data: nil,
		})
		return
	}

	var req models.UpdateProductCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}

	if err := h.service.Update(ctx.Request.Context(), oldPID, oldCID, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Relation updated successfully", 
		Data: nil,
	})
}

// Delete godoc
// @Summary Delete relation
// @Tags product-categories
// @Param product_id query int true "Product ID"
// @Param category_id query int true "Category ID"
// @Success 200 {object} models.WebResponse
// @Router /productcategory [delete]
func (h *ProductCategoryHandler) Delete(ctx *gin.Context) {
	pID, _ := strconv.Atoi(ctx.Query("product_id"))
	cID, _ := strconv.Atoi(ctx.Query("category_id"))

	if err := h.service.Delete(ctx.Request.Context(), pID, cID); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false, 
			Message: err.Error(), 
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true, 
		Message: "Relation deleted successfully", 
		Data: nil,
	})
}
