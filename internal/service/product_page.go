package service

import (
	"context"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductPageService struct {
	productRepo  *repository.ProductRepository
	discountRepo *repository.DiscountRepository
}

func NewProductPageService(pr *repository.ProductRepository, dr *repository.DiscountRepository) *ProductPageService {
	return &ProductPageService{
		productRepo:  pr,
		discountRepo: dr,
	}
}

func (s *ProductPageService) GetCatalogOnly(ctx context.Context, params map[string]string) (*models.ProductCatalogResponse, error) {
	return s.productRepo.GetCatalog(ctx, params)
}

func (s *ProductPageService) GetAllPromos(ctx context.Context) ([]models.Discount, error) {
	return s.discountRepo.FindAll(ctx)
}
