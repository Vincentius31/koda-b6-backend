package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"math/big"

	"github.com/matthewhartstonge/argon2"
)

type ForgotPasswordService struct {
	repoUser   *repository.UserRepository
	repoForgot *repository.ForgotPasswordRepository
}

func NewForgotPasswordRepositoryService(ru *repository.UserRepository, rf *repository.ForgotPasswordRepository) *ForgotPasswordService {
	return &ForgotPasswordService{
		repoUser: ru,
		repoForgot: rf,
	}
}

func (s *ForgotPasswordService) RequestForgotPassword(ctx context.Context, req models.ForgotPasswordRequest) error {
    _, err := s.repoUser.GetByEmail(ctx, req.Email)
    if err != nil {
        return err 
    }

    nBig, _ := rand.Int(rand.Reader, big.NewInt(900000))
    otpCode := int(nBig.Int64()) + 100000

	fmt.Printf("OTP Request %s: %d\n", req.Email, otpCode)

    return s.repoForgot.CreateForgetRequest(ctx, req.Email, otpCode)
}

func (s *ForgotPasswordService) ResetPassword(ctx context.Context, req models.ResetPasswordRequest) error {
    _, err := s.repoForgot.GetDataByEmailAndCode(ctx, req.Email, req.OTPCode)
    if err != nil {
        return err
    }

    user, err := s.repoUser.GetByEmail(ctx, req.Email)
    if err != nil {
        return err
    }

    argon := argon2.DefaultConfig()
    encoded, _ := argon.HashEncoded([]byte(req.NewPassword))
    user.Password = string(encoded)

    err = s.repoUser.Update(ctx, user.IDUser, *user)
    if err != nil {
        return err
    }

    return s.repoForgot.DeleteByCode(ctx, req.OTPCode)
}
