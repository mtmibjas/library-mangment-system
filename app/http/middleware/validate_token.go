package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	jwtSecretKey := []byte(jwtSecret)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		validate := func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing token"})
			}

			tokenString = tokenString[len("Bearer "):]

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				return jwtSecretKey, nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
			}

			userID, _ := claims["user_id"].(float64)
			role, _ := claims["role"].(float64)

			c.Set("user_id", userID)
			c.Set("role_id", role)

			return next(c)
		}
		return validate
	}
}
