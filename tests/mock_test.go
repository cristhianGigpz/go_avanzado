package tests

import (
	"go-avanzado/repositories"
	"testing"
)

func TestFindUser(t *testing.T) {

	mockRepo := repositories.MockUserRepository{}

	// service := AuthService{
	// 	repo: &mockRepo,
	// }

	// user, err := service.repo.FindByEmail(
	// 	"test@gmail.com",
	// )

	user, err := mockRepo.FindByEmail(
		"test@gmail.com",
	)

	if err != nil {
		t.Error(err)
	}

	if user.Email != "test@gmail.com" {
		t.Error("invalid email")
	}
}
