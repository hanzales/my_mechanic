package user

import (
	"MyMechanic/config"
	"github.com/AleksK1NG/api-mc/pkg/httpErrors"
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Auth handlers
type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	sessUC session.UCSession
	logger logger.Logger
}

// NewAuthHandlers Auth handlers constructor
func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, sessUC session.UCSession, log logger.Logger) auth.Handlers {
	return &authHandlers{cfg: cfg, authUC: authUC, sessUC: sessUC, logger: log}
}

// GetUserByID godoc
// @Summary get user by id
// @Description get string by ID
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param id path int true "user_id"
// @Success 200 {object} models.User
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/{id} [get]
func (h *authHandlers) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			//utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user, err := h.authUC.GetByID(ctx, uID)
		if err != nil {
			//utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, user)
	}
}
