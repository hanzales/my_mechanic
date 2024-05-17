package server

import (
	commentsRepository "MyMechanic/internal/comments/repository"
	apiMiddlewares "MyMechanic/middleware"
	"strings"

	"MyMechanic/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	// _ "github.com/AleksK1NG/api-mc/docs"
	commentsHttp "MyMechanic/internal/comments/delivery/http"
	commentsService "MyMechanic/internal/comments/service"
	usersHttp "MyMechanic/internal/users/delivery/http"

	usersRepository "MyMechanic/internal/users/repository"
	usersService "MyMechanic/internal/users/service"
)

// MapHandlers Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	commentsRepo := commentsRepository.CommentsRepository(s.db)
	usersRepo := usersRepository.UsersRepository(s.db)

	// Init services
	commentService := commentsService.CommentsService(s.cfg, commentsRepo, s.logger)
	userService := usersService.UsersService(s.cfg, usersRepo, s.logger)

	// Init handlers
	commHandlers := commentsHttp.CommentsHandlers(s.cfg, commentService, s.logger)
	userHandlers := usersHttp.UsersHandlers(s.cfg, userService, s.logger)

	//düzenlenecek. handler bazlı çalışması lazım
	mw := apiMiddlewares.NewMiddlewareManager(userService, s.cfg, []string{"*"}, s.logger)

	docs.SwaggerInfo.Title = "MyMechanic REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if s.cfg.Server.SSL {
		e.Pre(middleware.HTTPSRedirect())
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
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

	commentGroup := v1.Group("/comment")
	userGroup := v1.Group("/user")

	commentsHttp.MapCommentsRoutes(commentGroup, commHandlers, mw)
	usersHttp.MapUsersRoutes(userGroup, userHandlers)

	return nil
}
