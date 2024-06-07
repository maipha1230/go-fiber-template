package utils

import (
	"strings"
	"time"

	"example.com/prac02/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID uint, username string) (string, error) {
	config.LoadEnv()
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	config.LoadEnv()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(config.GetEnv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
	}

	return token, nil
}

func JWTMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ThrowExceoption(ctx, fiber.StatusUnauthorized, "Unauthorized")
	}

	tokenString := strings.Split(authHeader, " ")[1]
	token, err := ValidateJWT(tokenString)
	if err != nil {
		return ThrowExceoption(ctx, fiber.StatusUnauthorized, "Unauthorized")
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))
	ctx.Locals("userID", userID)

	return ctx.Next()
}
