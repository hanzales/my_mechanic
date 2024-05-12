package comments

import (
	"MyMechanic/internal/models"
	"context"
)

// Comments use case
type UseCase interface {
	GetByID(ctx context.Context, id int) (*models.Comment, error)
	Delete(ctx context.Context, id int) error
}
