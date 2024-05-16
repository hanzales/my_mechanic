package service

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/logger"
	"MyMechanic/pkg/utils"
	"context"
	"github.com/pkg/errors"
)

// Comments UseCase
type usersUC struct {
	cfg      *config.Config
	userRepo users.Repository
	logger   logger.Logger
}

// Users Service constructor
func UsersService(cfg *config.Config, userRepo users.Repository, logger logger.Logger) users.Service {
	return &usersUC{cfg: cfg, userRepo: userRepo, logger: logger}
}

func (u usersUC) Login(ctx context.Context, request *models.LoginRequest) (*models.UserWithToken, error) {
	foundUser, err := u.userRepo.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	if err = foundUser.ComparePasswords(request.Password); err != nil {
		return nil, models.NewUnauthorizedError(errors.Wrap(err, "usersService.GetUsers.ComparePasswords"))
	}

	token, err := utils.GenerateJWTToken(foundUser, u.cfg)
	if err != nil {
		return nil, models.NewInternalServerError(errors.Wrap(err, "usersService.GetUsers.GenerateJWTToken"))
	}

	userWithToken := &models.UserWithToken{User: foundUser, Token: token}
	return userWithToken, err
}

func (u usersUC) GetById(ctx context.Context, id int) (*models.User, error) {
	return u.userRepo.GetUserById(ctx, id)
}
