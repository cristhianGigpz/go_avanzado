package repositories

import "go-avanzado/models"

//Interface:
type UserRepository interface {
	FindByEmail(
		email string,
	) (*models.User, error)
}

//Servicio:
type AuthService struct {
	repo UserRepository
}

//Mock Repository
type MockUserRepository struct{}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {

	return &models.User{
		ID:    1,
		Email: email,
	}, nil
}
