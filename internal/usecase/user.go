package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/zsbahtiar/hello-echo-go/internal/entity"
)

type userUsecase struct{}

type UserUsecase interface {
	CreateUser(ctx context.Context, req *entity.User) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, req *entity.User) (*entity.User, error)
}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

func (u *userUsecase) CreateUser(ctx context.Context, req *entity.User) (*entity.User, error) {
	req.ID = uuid.NewString()
	return req, nil
}

func (u *userUsecase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	panic("implement me!")
}

func (u *userUsecase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	panic("implement me!")
}

func (u *userUsecase) DeleteUser(ctx context.Context, id string) error {
	panic("implement me!")
}

func (u *userUsecase) UpdateUser(ctx context.Context, req *entity.User) (*entity.User, error) {
	panic("implement me!")
}
