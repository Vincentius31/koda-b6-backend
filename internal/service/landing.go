package service

import (
	"context"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type LandingService struct {
	productRepo *repository.ProductRepository
	reviewRepo * repository.ReviewRepository
}

func NewLandingService(pr *repository.ProductRepository, rr *repository.ReviewRepository) *LandingService {
	return &LandingService{
		productRepo: pr,
		reviewRepo: rr,
	}
}

func (s *LandingService) GetRecommendedProducts(ctx context.Context) ([]models.ProductLanding, error) {
	return s.productRepo.GetRecommended(ctx)
}

func (s *LandingService) GetLatestReviews(ctx context.Context) ([]models.ReviewLanding, error) {
	return s.reviewRepo.GetLatestReviews(ctx)
}