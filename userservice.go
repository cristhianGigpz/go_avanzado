package main

type UserService struct{}

func (s *UserService) IsAdult(
	age int,
) bool {

	return age >= 18
}
