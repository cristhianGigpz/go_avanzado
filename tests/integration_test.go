package tests

import (
	"go-avanzado/models"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=gigpz dbname=bd_tests port=5434 sslmode=disable"

var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func TestCreateUser(
	t *testing.T,
) {
	t.Parallel()

	db.Exec("DELETE FROM users")

	user := models.User{
		ID:    22,
		Name:  "Enrique",
		Email: "enriqueo@gmail.com",
	}

	db.Create(&user)

	var count int64

	db.Model(&models.User{}).
		Count(&count)

	if count != 1 {
		t.Error("user not created")
	}
}
