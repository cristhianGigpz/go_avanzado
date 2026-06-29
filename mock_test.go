package main

import "testing"

func TestFindUser(t *testing.T) {

	mockRepo := MockUserRepository{}

	service := AuthService{
		repo: &mockRepo,
	}

	user, err := service.repo.FindByEmail(
		"test@gmail.com",
	)

	if err != nil {
		t.Error(err)
	}

	if user.Email != "test@gmail.com" {
		t.Error("invalid email")
	}
}
