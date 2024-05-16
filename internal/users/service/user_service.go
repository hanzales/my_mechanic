package service

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/logger"
	"context"
	"github.com/pkg/errors"
)

// Comments UseCase
type usersUC struct {
	cfg      *config.Config
	userRepo users.Repository
	logger   logger.Logger
}

// Comments UseCase constructor
func UsersService(cfg *config.Config, userRepo users.Repository, logger logger.Logger) users.Service {
	return &usersUC{cfg: cfg, userRepo: userRepo, logger: logger}
}

func (u usersUC) Login(ctx context.Context, request *models.LoginRequest) (*models.UserWithToken, error) {
	foundUser, err := u.userRepo.GetUserByEmail(ctx, request.Email)

	//if foundUser != nil || err == nil {
	//	return nil, models.NewRestErrorWithMessage(http.StatusBadRequest, models.ErrEmailAlreadyExists, nil)
	//}

	if err = foundUser.PrepareCreate(); err != nil {
		return nil, models.NewBadRequestError(errors.Wrap(err, ""))
	}

	if err = foundUser.ComparePasswords(request.Password); err != nil {
		return nil, models.NewUnauthorizedError(errors.Wrap(err, "usersservice.GetUsers.ComparePasswords"))
	}

	return nil, err
}
