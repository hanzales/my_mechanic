package http

import (
	"github.com/labstack/echo/v4"

	"MyMechanic/internal/comments"
)

// Map comments routes
func MapCommentsRoutes(commGroup *echo.Group, h comments.Handlers) {
	commGroup.POST("", h.Create())
	commGroup.DELETE("/:comment_id", h.Delete())
	commGroup.PUT("/:comment_id", h.Update())
	commGroup.GET("/:comment_id", h.GetByID())
	commGroup.GET("/byNewsId/:news_id", h.GetAllByNewsID())
}
