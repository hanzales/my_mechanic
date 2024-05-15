package http

import (
	"MyMechanic/internal/users"
	"github.com/labstack/echo/v4"
)

// Map comments routes
func MapUsersRoutes(userGroup *echo.Group, h users.Handlers) {
	userGroup.GET("/get", h.GetByID())
	userGroup.POST("/login", h.Login())
}
