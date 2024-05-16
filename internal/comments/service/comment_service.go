package service

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
func CommentsService(cfg *config.Config, commRepo comments.Repository, logger logger.Logger) comments.Service {
	return &commentsUC{cfg: cfg, commRepo: commRepo, logger: logger}
}

// GetByID comment
func (u *commentsUC) GetByID(ctx context.Context, id int) (*models.Comment, error) {
	return u.commRepo.GetByID(ctx, id)
}

func (u *commentsUC) Delete(ctx context.Context, id int) error {
	return u.commRepo.Delete(ctx, id)
}

func (u *commentsUC) Create(ctx context.Context, addCommentRequest *models.AddCommentRequest) (*models.Comment, error) {
	return u.commRepo.Create(ctx, addCommentRequest)
}

func (u *commentsUC) Update(ctx context.Context, updateCommentRequest *models.UpdateCommentRequest) (*models.Comment, error) {
	return u.commRepo.Update(ctx, updateCommentRequest)
}

func (u *commentsUC) IncreaseLikeCount(ctx context.Context, increaseLikeRequest *models.IncreaseLikeRequest) error {
	return u.commRepo.IncreaseLikeCount(ctx, increaseLikeRequest)
}
