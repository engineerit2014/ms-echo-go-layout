package repositories

import (
	"context"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/entities"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user entities.User) error
	GetByEmail(ctx context.Context, email string) (entities.User, error)
	UpdateByEmail(ctx context.Context, user entities.User, email string) error
	DeleteByEmail(ctx context.Context, email string) error
}
