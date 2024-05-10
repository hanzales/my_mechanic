package server

import (
	"strings"

	"github.com/AleksK1NG/api-mc/docs"
	"github.com/AleksK1NG/api-mc/pkg/csrf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	// _ "github.com/AleksK1NG/api-mc/docs"
	commentsHttp "MyMechanic/internal/comments/delivery/http"
	commentsRepository "MyMechanic/internal/comments/repository"
	commentsUseCase "MyMechanic/internal/comments/usecase"
)

// MapHandlers Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	cRepo := commentsRepository.NewCommentsRepository(s.db)

	// Init useCases
	commUC := commentsUseCase.NewCommentsUseCase(s.cfg, cRepo, s.logger)

	// Init handlers
	commHandlers := commentsHttp.NewCommentsHandlers(s.cfg, commUC, s.logger)

	docs.SwaggerInfo.Title = "Go example REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if s.cfg.Server.SSL {
		e.Pre(middleware.HTTPSRedirect())
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	v1 := e.Group("/api/v1")

	commGroup := v1.Group("/comments")

	commentsHttp.MapCommentsRoutes(commGroup, commHandlers)

	return nil
}
