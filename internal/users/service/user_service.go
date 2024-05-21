package service

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/logger"
	"MyMechanic/pkg/utils"
	"context"
	"github.com/pkg/errors"
	"net/http"
	"time"
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

func (u usersUC) Register(ctx context.Context, request *models.RegisterRequest) (*models.UserWithToken, error) {
	foundUser, err := u.userRepo.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	//Maile kayıtlı kullanıcı varsa hata ver
	if foundUser != nil {
		return nil, models.NewRestErrorWithMessage(http.StatusBadRequest, models.ErrEmailAlreadyExists, nil)
	}

	user := &models.User{0,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Password,
		request.Role,
		request.About,
		request.Avatar,
		request.PhoneNumber,
		request.Address,
		request.City,
		request.Country,
		request.Gender,
		request.Postcode,
		request.Birthday,
		true,
		time.Now(),
		time.Now(),
		time.Now()}

	if err = user.PrepareCreate(); err != nil {
		return nil, models.NewBadRequestError(errors.Wrap(err, "authUC.Register.PrepareCreate"))
	}

	request.Password = user.Password
	createdUser, err := u.userRepo.Register(ctx, request)
	if err != nil {
		return nil, err
	}
	createdUser.SanitizePassword()

	token, err := utils.GenerateJWTToken(createdUser, u.cfg)
	if err != nil {
		return nil, models.NewInternalServerError(errors.Wrap(err, "authUC.Register.GenerateJWTToken"))
	}

	userWithToken := &models.UserWithToken{User: createdUser, Token: token}
	return userWithToken, err
}
