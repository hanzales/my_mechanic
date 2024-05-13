package comments

import (
	"MyMechanic/internal/models"
	"context"
)

// Comments use case
type UseCase interface {
	GetByID(ctx context.Context, id int) (*models.Comment, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, addCommentRequest *models.AddCommentRequest) (*models.Comment, error)
	Update(ctx context.Context, updateCommentRequest *models.UpdateCommentRequest) (*models.Comment, error)
	IncreaseLikeCount(ctx context.Context, increaseLikeRequest *models.IncreaseLikeRequest) error
}
