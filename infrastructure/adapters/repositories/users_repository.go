package repository

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	respKit "github.com/laironacosta/kit-go/middleware/responses"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/entities"
	"github.com/laironacosta/ms-echo-go-layout/domain/users/ports/repositories"
	"github.com/laironacosta/ms-echo-go-layout/infrastructure/enums"
	"strings"
)

type userRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) repositories.UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

func (r *userRepository) Create(ctx context.Context, user entities.User) error {
	_, err := r.db.Model(&user).Context(ctx).Insert()
	if err != nil {
		if pgErr := err.(pg.Error); pgErr != nil && strings.Contains(pgErr.Error(), enums.ErrorDBDuplicatedKeyMsg) {
			return respKit.GenericAlreadyExistsError(enums.ErrorEmailExistsCode, fmt.Sprintf(enums.ErrorEmailExistsMsg, user.Email))
		}

		return respKit.GenericBadRequestError(enums.ErrorInsertCode, err.Error())
	}

	return nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (entities.User, error) {
	u := entities.User{}
	err := r.db.Model(&u).Context(ctx).Where("email = ?", email).Select()
	if err != nil {
		switch err {
		case pg.ErrNoRows:
			return u, respKit.GenericNotFoundError(enums.ErrorEmailNotFoundCode, fmt.Sprintf(enums.ErrorEmailNotFoundMsg, email))
		default:
			return u, respKit.GenericBadRequestError(enums.ErrorGetByEmailCode, err.Error())
		}
	}

	return u, nil
}

func (r *userRepository) UpdateByEmail(ctx context.Context, user entities.User, email string) error {
	err := r.db.Model(&entities.User{}).Context(ctx).Where("email = ?", email).Select()
	if err != nil && err == pg.ErrNoRows {
		return respKit.GenericNotFoundError(enums.ErrorEmailNotFoundCode, fmt.Sprintf(enums.ErrorEmailNotFoundMsg, email))
	}

	if _, err := r.db.Model(&entities.User{}).Context(ctx).Set("name = ?", user.Name).Where("email = ?", email).Update(); err != nil {
		return respKit.GenericBadRequestError(enums.ErrorUpdateCode, err.Error())
	}

	return nil
}

func (r *userRepository) DeleteByEmail(ctx context.Context, email string) error {
	if _, err := r.db.Model(&entities.User{}).Context(ctx).Where("email = ?", email).Delete(); err != nil {
		switch err {
		case pg.ErrNoRows:
			return respKit.GenericNotFoundError(enums.ErrorEmailNotFoundCode, fmt.Sprintf(enums.ErrorEmailNotFoundMsg, email))
		default:
			return respKit.GenericBadRequestError(enums.ErrorDeleteCode, err.Error())
		}
	}

	return nil
}
