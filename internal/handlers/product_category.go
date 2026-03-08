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

func (h *ProductCategoryHandler) Create(ctx *gin.Context) {
	var req models.CreateProductCategoryRequest
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
		Message: "Relation created",
	})
}

func (h *ProductCategoryHandler) GetAll(ctx *gin.Context) {
	data, err := h.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Data:    data,
	})
}

func (h *ProductCategoryHandler) GetByID(ctx *gin.Context) {
	pID, _ := strconv.Atoi(ctx.Query("product_id"))
	cID, _ := strconv.Atoi(ctx.Query("category_id"))

	data, err := h.service.GetByID(ctx.Request.Context(), pID, cID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.WebResponse{
			Success: false,
			Message: "Relation not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Data:    data,
	})
}

func (h *ProductCategoryHandler) Update(ctx *gin.Context) {
	oldPID, errP := strconv.Atoi(ctx.Query("old_product_id"))
	oldCID, errC := strconv.Atoi(ctx.Query("old_category_id"))

	if errP != nil || errC != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: "Old Product ID and Old Category ID are required",
		})
		return
	}

	var req models.UpdateProductCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	
	if err := h.service.Update(ctx.Request.Context(), oldPID, oldCID, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Relation updated successfully",
	})
}

func (h *ProductCategoryHandler) Delete(ctx *gin.Context) {
	pID, _ := strconv.Atoi(ctx.Query("product_id"))
	cID, _ := strconv.Atoi(ctx.Query("category_id"))

	if err := h.service.Delete(ctx.Request.Context(), pID, cID); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.WebResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.WebResponse{
		Success: true,
		Message: "Relation deleted",
	})
}
