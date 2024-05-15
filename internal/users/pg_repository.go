package users

import (
	"MyMechanic/internal/models"
	"context"
)

// Comments repository interface
type Repository interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, request models.LoginRequest) (*models.User, error)
}