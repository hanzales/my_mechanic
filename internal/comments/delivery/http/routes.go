package http

import (
	"MyMechanic/middleware"
	"github.com/labstack/echo/v4"

	"MyMechanic/internal/comments"
)

// Map comments routes
func MapCommentsRoutes(commGroup *echo.Group, h comments.Handlers, mw *middleware.MiddlewareManager) {
	commGroup.GET("/get", h.GetByID())
	commGroup.DELETE("/:id", h.Delete(), mw.AuthJWTMiddleware())
	commGroup.POST("/add", h.Create(), mw.AuthJWTMiddleware())
	commGroup.PUT("/update", h.Update(), mw.AuthJWTMiddleware())
	commGroup.POST("/increase-like-count", h.IncreaseLikeCount(), mw.AuthJWTMiddleware())
}
