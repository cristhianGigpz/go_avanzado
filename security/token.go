package security

import (
	"go-avanzado/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_key_gigpz")

func GenerateJWT(user models.User) (string, error) {

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp": time.Now().
			Add(time.Minute * 15).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {

			return jwtSecret, nil
		},
	)
}
