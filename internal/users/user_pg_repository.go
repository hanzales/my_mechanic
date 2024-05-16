package users

import (
	"MyMechanic/internal/models"
	"context"
)

// Comments repository interface
type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
