package services

import (
	"go-avanzado/models"
	"go-avanzado/security"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context, db *gorm.DB) {

	var body struct {
		Name     string
		Email    string
		Age      int
		Password string
	}

	c.BindJSON(&body)

	hashedPassword, err := security.HashPassword(
		body.Password,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Age:      body.Age,
		Password: hashedPassword,
	}

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Error creating user",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Usuario creado",
	})
}

func LoginUser(c *gin.Context, db *gorm.DB) {

	var body struct {
		Email    string
		Password string
	}

	c.BindJSON(&body)

	var user models.User

	err := db.Where(
		"email = ?",
		body.Email,
	).First(&user).Error

	if err != nil {

		c.JSON(401, gin.H{
			"error": "Credenciales inválidas",
		})

		return
	}

	valid := security.CheckPassword(
		body.Password,
		user.Password,
	)

	if !valid {

		c.JSON(401, gin.H{
			"error": "Credenciales inválidas",
		})

		return
	}

	token, err := security.GenerateJWT(models.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})

	if err != nil {

		c.JSON(500, gin.H{
			"error": "Error generando token",
		})

		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
