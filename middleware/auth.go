package middleware

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"MyMechanic/internal/users"
	"MyMechanic/pkg/utils"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

// JWT way of auth using cookie or Authorization header
func (mw *MiddlewareManager) AuthJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerHeader := c.Request().Header.Get("Authorization")

			mw.logger.Infof("auth middleware bearerHeader %s", bearerHeader)

			if bearerHeader != "" {
				headerParts := strings.Split(bearerHeader, " ")
				if len(headerParts) != 2 {
					mw.logger.Error("auth middleware", zap.String("headerParts", "len(headerParts) != 2"))
					return c.JSON(http.StatusUnauthorized, models.NewUnauthorizedError(models.Unauthorized))
				}

				tokenString := headerParts[1]

				if err := mw.validateJWTToken(tokenString, mw.usersService, c, mw.cfg); err != nil {
					mw.logger.Error("middleware validateJWTToken", zap.String("headerJWT", err.Error()))
					return c.JSON(http.StatusUnauthorized, models.NewUnauthorizedError(models.Unauthorized))
				}

				return next(c)
			}

			cookie, err := c.Cookie("jwt-token")
			if err != nil {
				mw.logger.Errorf("c.Cookie", err.Error())
				return c.JSON(http.StatusUnauthorized, models.NewUnauthorizedError(models.Unauthorized))
			}

			if err = mw.validateJWTToken(cookie.Value, mw.usersService, c, mw.cfg); err != nil {
				mw.logger.Errorf("validateJWTToken", err.Error())
				return c.JSON(http.StatusUnauthorized, models.NewUnauthorizedError(models.Unauthorized))
			}

			return next(c)
		}
	}
}

func (mw *MiddlewareManager) validateJWTToken(tokenString string, userService users.Service, c echo.Context, cfg *config.Config) error {
	if tokenString == "" {
		return models.InvalidJWTToken
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
		}
		secret := []byte(cfg.Server.JwtSecretKey)
		return secret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return models.InvalidJWTToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdStr, ok := claims["user_id"].(string)
		if !ok {
			return models.InvalidJWTClaims
		}

		userId, err := strconv.Atoi(userIdStr)

		if err != nil {
			return err
		}

		u, err := userService.GetById(c.Request().Context(), userId)
		if err != nil {
			return err
		}

		c.Set("user", u)

		ctx := context.WithValue(c.Request().Context(), utils.UserCtxKey{}, u)
		c.SetRequest(c.Request().WithContext(ctx))
	}
	return nil
}
