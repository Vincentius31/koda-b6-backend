package handlers

import (
    "koda-b6-backend/internal/models"
    "koda-b6-backend/internal/service"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type DashboardHandler struct {
    dashboardService *service.DashboardService
}

func NewDashboardHandler(ds *service.DashboardService) *DashboardHandler {
    return &DashboardHandler{dashboardService: ds}
}

func (h *DashboardHandler) GetSalesCategory(ctx *gin.Context) {
    data, err := h.dashboardService.GetSalesByCategory(ctx.Request.Context())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, models.WebResponse{
            Success: false,
            Message: "Failed to fetch sales category: " + err.Error(),
            Data:    nil,
        })
        return
    }

    ctx.JSON(http.StatusOK, models.WebResponse{
        Success: true,
        Message: "Success fetching sales by category",
        Data:    data,
    })
}

func (h *DashboardHandler) GetBestSellers(ctx *gin.Context) {
    limitStr := ctx.DefaultQuery("limit", "10")
    limit, err := strconv.Atoi(limitStr)
    if err != nil {
        limit = 10
    }

    data, err := h.dashboardService.GetBestSellers(ctx.Request.Context(), limit)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, models.WebResponse{
            Success: false,
            Message: "Failed to fetch best sellers: " + err.Error(),
            Data:    nil,
        })
        return
    }

    ctx.JSON(http.StatusOK, models.WebResponse{
        Success: true,
        Message: "Success fetching best sellers",
        Data:    data,
    })
}

func (h *DashboardHandler) GetOrderStats(ctx *gin.Context) {
    data, err := h.dashboardService.GetOrderStats(ctx.Request.Context())
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
        Message: "Success fetching order stats",
        Data:    data,
    })
}