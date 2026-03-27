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

	params := map[string]string{"page": "1", "limit": "3"}
	recommended, err := s.productRepo.GetCatalog(ctx, params)
	if err != nil {
		return nil, err
	}

	var finalRecommended []models.ProductCatalog
	for _, item := range recommended.Items {
		if item.IDProduct != id {
			finalRecommended = append(finalRecommended, item)
		}
	}

	return &models.ProductDetailResponse{
		Product:     *detail,
		Recommended: finalRecommended,
	}, nil
}