package service

import (
	"context"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type LandingService struct {
	productRepo *repository.ProductRepository
}

func NewLandingService(pr *repository.ProductRepository) *LandingService {
	return &LandingService{
		productRepo: pr,
	}
}

func (s *LandingService) GetRecommendedProducts(ctx context.Context) ([]models.ProductLanding, error) {
	return s.productRepo.GetRecommended(ctx)
}