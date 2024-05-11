package usecase

import (
	"MyMechanic/config"
	"MyMechanic/internal/comments"
	"MyMechanic/internal/models"
	"MyMechanic/pkg/logger"
	"context"
)

// Comments UseCase
type commentsUC struct {
	cfg      *config.Config
	commRepo comments.Repository
	logger   logger.Logger
}

// Comments UseCase constructor
func NewCommentsUseCase(cfg *config.Config, commRepo comments.Repository, logger logger.Logger) comments.UseCase {
	return &commentsUC{cfg: cfg, commRepo: commRepo, logger: logger}
}

// GetByID comment
func (u *commentsUC) GetByID(ctx context.Context, id int) (*models.Comment, error) {
	return u.commRepo.GetByID(ctx, id)
}
