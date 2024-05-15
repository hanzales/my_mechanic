package usecase

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/logger"
	"context"
)

// Comments UseCase
type usersUC struct {
	cfg      *config.Config
	userRepo users.Repository
	logger   logger.Logger
}

// Comments UseCase constructor
func NewUsersUseCase(cfg *config.Config, userRepo users.Repository, logger logger.Logger) users.UseCase {
	return &usersUC{cfg: cfg, userRepo: userRepo, logger: logger}
}

func (u usersUC) GetByID(ctx context.Context, id int) (*models.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u usersUC) Login(ctx context.Context, request models.LoginRequest) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
