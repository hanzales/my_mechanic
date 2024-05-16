package users

import (
	"MyMechanic/internal/models"
	"context"
)

type Service interface {
	Login(ctx context.Context, request *models.LoginRequest) (*models.UserWithToken, error)
}
