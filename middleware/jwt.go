package middleware

import (
	"context"
	"go-avanzado/security"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func JWTMiddleware(rdb *redis.Client, ctx context.Context) gin.HandlerFunc {

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

		_, errT := rdb.Get(
			ctx,
			tokenString,
		).Result()

		if errT == nil {

			c.JSON(401, gin.H{
				"error": "token invalidado",
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
