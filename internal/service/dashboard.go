package service

import (
	"context"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type DashboardService struct {
	repo *repository.DashboardRepository
}

func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetSalesByCategory(ctx context.Context) ([]models.SalesCategory, error) {
	return s.repo.GetSalesByCategory(ctx)
}

func (s *DashboardService) GetBestSellers(ctx context.Context, limit int) ([]models.BestSeller, error) {
	if limit <= 0 {
		limit = 10
	}
	return s.repo.GetBestSellers(ctx, limit)
}

func (s *DashboardService) GetOrderStats(ctx context.Context) (*models.OrderStats, error) {
	return s.repo.GetOrderStats(ctx)
}
