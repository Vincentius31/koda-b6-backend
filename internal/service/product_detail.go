package service

import (
	"context"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type DetailProductService struct {
	productRepo *repository.ProductRepository
	db          *repository.ProductRepository
}

func NewDetailProductService(pr *repository.ProductRepository) *DetailProductService {
	return &DetailProductService{productRepo: pr}
}

func (s *DetailProductService) GetDetailByID(ctx context.Context, id int) (*models.ProductDetailResponse, error) {
	detail, err := s.productRepo.GetFullDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	recommended, err := s.productRepo.GetRandomRecommended(ctx, id, 15)
	if err != nil {
		recommended = []models.ProductCatalog{} 
	}

	return &models.ProductDetailResponse{
		Product:     *detail,
		Recommended: recommended,
	}, nil
}