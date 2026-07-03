package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

func RateLimiterMiddleware(rdb *redis.Client, ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip

		// 1. Incrementamos el contador
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error de servidor"})
			c.Abort()
			return
		}
		//Expirar
		// 2. Si es 1, garantizamos el TTL. Si por alguna razón no tiene TTL (devuelve -1), también lo fijamos.
		if count == 1 {
			rdb.Expire(ctx, key, time.Minute)
		} else {
			ttl, _ := rdb.TTL(ctx, key).Result()
			if ttl == -1 {
				rdb.Expire(ctx, key, time.Minute)
			}
		}
		// Bloquear
		// 3. Validamos el límite (Máximo 100 peticiones por minuto)
		if count > 100 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "demasiadas peticiones, intenta más tarde",
			})
			c.Abort() // Detiene la ejecución de los siguientes handlers
			return
		}

		c.Next()
	}
}
