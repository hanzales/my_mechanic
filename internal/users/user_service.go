package users

import (
	"MyMechanic/internal/models"
	"context"
)

type Service interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, request models.LoginRequest) (*models.UserWithToken, error)
}
