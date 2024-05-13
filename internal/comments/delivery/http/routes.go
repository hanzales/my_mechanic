package http

import (
	"github.com/labstack/echo/v4"

	"MyMechanic/internal/comments"
)

// Map comments routes
func MapCommentsRoutes(commGroup *echo.Group, h comments.Handlers) {
	commGroup.GET("/get", h.GetByID())
	commGroup.DELETE("/:id", h.Delete())
	commGroup.POST("/add", h.Create())
	commGroup.PUT("/update", h.Update())
	commGroup.POST("/increase-like-count", h.IncreaseLikeCount())
}
