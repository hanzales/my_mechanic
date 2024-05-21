package users

import (
	"MyMechanic/internal/models"
	"context"
)

type Service interface {
	Login(ctx context.Context, request *models.LoginRequest) (*models.UserWithToken, error)
	GetById(ctx context.Context, id int) (*models.User, error)
	Register(ctx context.Context, request *models.RegisterRequest) (*models.UserWithToken, error)
}
