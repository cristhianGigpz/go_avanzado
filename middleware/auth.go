package middleware

import "github.com/gin-gonic/gin"

func HeadersMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header(
			"X-API-VERSION",
			"1.0",
		)

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		if token == "" {

			c.JSON(401, gin.H{
				"error": "Token requerido",
			})

			c.Abort()

			return
		}

		c.Set("userID", "12345")
		c.Next()
	}
}
