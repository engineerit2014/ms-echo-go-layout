package users

import (
	"context"
	respKit "github.com/laironacosta/kit-go/middleware/responses"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/dto"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/entities"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/enums"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/ports/repositories"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/ports/services"
	"strings"
)

type userService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) services.UserServiceInterface {
	return &userService{
		userRepo,
	}
}

func (s *userService) Create(ctx context.Context, request dto.CreateUserRequest) error {
	user := entities.User{
		Name:  request.Name,
		Email: request.Email,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return dto.UserResponse{}, respKit.GenericBadRequestError(enums.ErrorEmailNotEmptyCode, enums.ErrorEmailNotEmptyMsg)
	}

	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *userService) UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return respKit.GenericBadRequestError(enums.ErrorEmailNotEmptyCode, enums.ErrorEmailNotEmptyMsg)
	}

	user := entities.User{
		Name:  request.Name,
		Email: email,
	}

	if err := s.userRepo.UpdateByEmail(ctx, user, email); err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteByEmail(ctx context.Context, email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return respKit.GenericBadRequestError(enums.ErrorEmailNotEmptyCode, enums.ErrorEmailNotEmptyMsg)
	}

	err := s.userRepo.DeleteByEmail(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
