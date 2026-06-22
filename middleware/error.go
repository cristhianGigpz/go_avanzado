package middleware

import "github.com/gin-gonic/gin"

func ErrorMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		errors := c.Errors

		if len(errors) > 0 {

			c.JSON(500, gin.H{
				"error": errors[0].Error(),
			})
		}
	}
}
