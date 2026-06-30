package tests

import (
	"go-avanzado/services"
	"testing"
)

func TestIsAdult(t *testing.T) {

	service := services.UserService{}

	result := service.IsAdult(20)

	if !result {
		t.Error("expected true")
	}
}
