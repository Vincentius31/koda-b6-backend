package service

import (
	"context"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) FindByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Register(ctx context.Context, req models.CreateUserRequest) error {
	newUser := models.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: req.Password,
	}
	return s.repo.Create(ctx, newUser)
}

func (s *UserService) Update(ctx context.Context, id int, req models.UpdateUserRequest) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Address != "" {
		user.Address = &req.Address
	}
	if req.Phone != "" {
		user.Phone = &req.Phone
	}
	if req.ProfilePicture != "" {
		user.ProfilePicture = &req.ProfilePicture
	}

	return s.repo.Update(ctx, id, *user)
}

func (s *UserService) Remove(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}