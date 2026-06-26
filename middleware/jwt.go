package middleware

import (
	"go-avanzado/security"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader(
			"Authorization",
		)

		if authHeader == "" {

			c.JSON(401, gin.H{
				"error": "Token requerido",
			})

			c.Abort()

			return
		}

		tokenString := strings.Replace(
			authHeader,
			"Bearer ",
			"",
			1,
		)

		token, err := security.ValidateJWT(
			tokenString,
		)

		if err != nil || !token.Valid {

			c.JSON(401, gin.H{
				"error": "Token inválido",
			})

			c.Abort()

			return
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set(
			"user_id",
			claims["user_id"],
		)

		c.Set("role",
			claims["role"],
		)

		c.Next()
	}
}
