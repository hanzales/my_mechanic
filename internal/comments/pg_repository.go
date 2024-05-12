package comments

import (
	"MyMechanic/internal/models"
	"context"
)

// Comments repository interface
type Repository interface {
	GetByID(ctx context.Context, id int) (*models.Comment, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, addCommentRequest *models.AddCommentRequest) (*models.AddCommentRequest, error)
}
