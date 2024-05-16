package users

import "github.com/labstack/echo/v4"

// Users HTTP Handlers interface
type Handlers interface {
	GetByID() echo.HandlerFunc
	Login() echo.HandlerFunc
}
