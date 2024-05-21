package http

import (
	"MyMechanic/internal/users"
	"github.com/labstack/echo/v4"
)

// Map comments routes
func MapUsersRoutes(userGroup *echo.Group, h users.Handlers) {
	userGroup.POST("/login", h.Login())
	userGroup.POST("/register", h.Register())
}
