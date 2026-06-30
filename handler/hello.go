package handler

import (
	"go-avanzado/models"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {

	c.JSON(200, models.User{
		ID:    1,
		Email: "juan@gmail.com",
		Name:  "Juan",
		Role:  "user",
	})
}
