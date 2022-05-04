package services

import (
	"context"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/dto"
)

type UserServiceInterface interface {
	Create(ctx context.Context, request dto.CreateUserRequest) error
	GetByEmail(ctx context.Context, email string) (dto.UserResponse, error)
	UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error
	DeleteByEmail(ctx context.Context, email string) error
}
