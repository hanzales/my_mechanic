package middleware

import (
	"MyMechanic/config"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	usersService users.Service
	cfg          *config.Config
	origins      []string
	logger       logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(usersService users.Service, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{usersService: usersService, cfg: cfg, origins: origins, logger: logger}
}
