package utils

import (
	"MyMechanic/config"
	"MyMechanic/internal/models"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// JWT Claims struct
type Claims struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

// Generate new JWT Token
func GenerateJWTToken(user *models.User, config *config.Config) (string, error) {
	// Register the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email:  user.Email,
		UserId: strconv.Itoa(user.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Register the JWT string
	tokenString, err := token.SignedString([]byte(config.Server.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
