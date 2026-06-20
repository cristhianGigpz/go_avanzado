package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Antes del handler")
		c.Next()
		println("Despues del handler")
	}
}

func LoggerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		duration := time.Since(start)

		fmt.Println(
			c.Request.Method,
			c.Request.URL.Path,
			duration,
		)
	}
}

func Init() {
	r := gin.New()

	//r.Use(MyMiddleware())
	r.Use(LoggerMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Go Avanzado !",
		})
	})

	r.GET("/users", MyMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Esta es la ruta para ver los usuarios",
		})
	})

	r.Run(":8080")
}
