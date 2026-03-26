package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func validateUser(fullname string, email string, password string) error {
	if len(strings.TrimSpace(fullname)) < 1 {
		return errors.New("Fullname must be at least 1 characters")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return errors.New("Invalid email format")
	}
	if strings.Index(email, "@") > strings.Index(email, ".") {
		return errors.New("Invalid email domain format")
	}
	if len(password) < 5 {
		return errors.New("Password must be at least 5 characters")
	}
	return nil
}

func (s *UserService) Login(ctx context.Context, req models.LoginRequest) (string, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)

	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	ok, err := argon2.VerifyEncoded([]byte(req.Password), []byte(user.Password))
	if err != nil || !ok {
		return "", errors.New("Invalid Email or Password")
	}

	claims := jwt.MapClaims{
		"user_id": user.IDUser,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Minute * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("APP_SECRET")
	return token.SignedString([]byte(secret))
}

func (s *UserService) UpdateProfileImage(ctx context.Context, id int, filename string) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	user.ProfilePicture = &filename
	return s.repo.Update(ctx, id, *user)
}

func (s *UserService) FindAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) FindByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Register(ctx context.Context, req models.CreateUserRequest) error {
	if err := validateUser(req.Fullname, req.Email, req.Password); err != nil {
		return err
	}

	existingUser, _ := s.repo.GetByEmail(ctx, req.Email)

	if existingUser != nil {
		return errors.New("Email is already registered!")
	}

	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(req.Password))

	if err != nil {
		return err
	}

	newUser := models.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: string(encoded),
	}
	return s.repo.Create(ctx, newUser)
}

func (s *UserService) Update(ctx context.Context, id int, req models.UpdateUserRequest) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if req.Fullname != "" {
		if len(strings.TrimSpace(req.Fullname)) < 1 {
			return errors.New("Fullname cannot be empty!")
		}
		user.Fullname = req.Fullname
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		if len(req.Password) < 5 {
			return errors.New("password must be at least 5 characters")
		}
		argon := argon2.DefaultConfig()
		encoded, _ := argon.HashEncoded([]byte(req.Password))
		user.Password = string(encoded)
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
