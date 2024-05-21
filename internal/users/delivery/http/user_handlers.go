package http

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/logger"
	"MyMechanic/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Users handlers
type usersHandlers struct {
	cfg    *config.Config
	userUC users.Service
	logger logger.Logger
}

// UsersHandlers Userss handlers constructor
func UsersHandlers(cfg *config.Config, userUC users.Service, logger logger.Logger) users.Handlers {
	return &usersHandlers{cfg: cfg, userUC: userUC, logger: logger}
}

func (u usersHandlers) Login() echo.HandlerFunc {
	return func(c echo.Context) error {

		loginRequest := &models.LoginRequest{}
		err := utils.SanitizeRequest(c, loginRequest)

		if err != nil {
			return utils.ErrResponseWithLog(c, u.logger, err)
		}

		userWithToken, err := u.userUC.Login(c.Request().Context(), loginRequest)

		if err != nil {
			utils.LogResponseError(c, u.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(userWithToken))
	}
}

func (u usersHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		registerRequest := &models.RegisterRequest{}
		err := utils.SanitizeRequest(c, registerRequest)

		if err != nil {
			return utils.ErrResponseWithLog(c, u.logger, err)
		}

		userWithToken, err := u.userUC.Register(c.Request().Context(), registerRequest)

		if err != nil {
			utils.LogResponseError(c, u.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(userWithToken))
	}
}
