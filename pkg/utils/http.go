package utils

import (
	"MyMechanic/internal/models"
	"MyMechanic/pkg/logger"
	"MyMechanic/pkg/sanitize"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

// UserCtxKey is a key used for the User object in the context
type UserCtxKey struct{}

// Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

// Error response with logging error for echo context
func ErrResponseWithLog(ctx echo.Context, logger logger.Logger, err error) error {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetIPAddress(ctx),
		err,
	)
	return ctx.JSON(models.ErrorResponse(err))
}

// Error response with logging error for echo context
func LogResponseError(ctx echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetIPAddress(ctx),
		err,
	)
}

// Read sanitize and validate request
func SanitizeRequest(ctx echo.Context, request interface{}) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	defer ctx.Request().Body.Close()

	sanBody, err := sanitize.SanitizeJSON(body)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	if err = json.Unmarshal(sanBody, request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request().Context(), request)
}
