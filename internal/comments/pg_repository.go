package comments

import (
	"MyMechanic/internal/models"
	"context"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/google/uuid"
)

// Comments repository interface
type Repository interface {
	Create(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	Update(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	Delete(ctx context.Context, commentID uuid.UUID) error
	GetByID(ctx context.Context, commentID uuid.UUID) (*models.CommentBase, error)
	GetAllByNewsID(ctx context.Context, newsID uuid.UUID, query *utils.PaginationQuery) (*models.CommentsList, error)
}
