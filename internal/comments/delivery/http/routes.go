package http

import (
	"github.com/labstack/echo/v4"

	"MyMechanic/internal/comments"
)

// Map comments routes
func MapCommentsRoutes(commGroup *echo.Group, h comments.Handlers) {
	commGroup.GET("/:id", h.GetByID())
}
