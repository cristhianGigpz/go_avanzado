package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5)

func RateLimitMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if !limiter.Allow() {

			c.JSON(429, gin.H{
				"error": "Demasiadas peticiones",
			})

			c.Abort()

			return
		}

		c.Next()
	}
}
