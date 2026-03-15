package service

import (
	"context"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductPageService struct {
	productRepo *repository.ProductRepository
}

func NewProductPageService(pr *repository.ProductRepository) *ProductPageService {
	return &ProductPageService{
		productRepo: pr,
	}
}

func (s *ProductPageService) GetCatalogOnly(ctx context.Context, params map[string]string) (*models.ProductCatalogResponse, error) {
	return s.productRepo.GetCatalog(ctx, params)
}
