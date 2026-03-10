package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(ctx context.Context, req models.CreateProductRequest) error {
	if req.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	product := models.Product{
		Name:     req.Name,
		Desc:     req.Desc,
		Price:    req.Price,
		Quantity: req.Quantity,
		IsActive: req.IsActive,
	}
	return s.repo.Create(ctx, product)
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id int) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, id int, req models.UpdateProductRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Product not found")
	}

	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Desc != nil {
		existing.Desc = *req.Desc
	}
	if req.Price != nil {
		if *req.Price < 0 {
			return errors.New("Price cannot be negative")
		}
		existing.Price = *req.Price
	}
	if req.Quantity != nil {
		if *req.Quantity < 0 {
			return errors.New("Quantity cannot be negative")
		}
		existing.Quantity = *req.Quantity
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
