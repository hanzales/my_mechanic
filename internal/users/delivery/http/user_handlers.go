package http

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/logger"
	"MyMechanic/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func (u usersHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			utils.LogResponseError(c, u.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		user, err := u.userUC.GetByID(c.Request().Context(), userId)
		if err != nil {
			utils.LogResponseError(c, u.logger, err)
			return c.JSON(models.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, models.NewSuccessResponse(user))
	}
}

func (u usersHandlers) Login() echo.HandlerFunc {
	return nil
}
