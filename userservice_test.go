package main

import "testing"

func TestIsAdult(t *testing.T) {

	service := UserService{}

	result := service.IsAdult(20)

	if !result {
		t.Error("expected true")
	}
}
