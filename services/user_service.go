package services

import (
	"maintenance-tracker/models"
	"maintenance-tracker/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) CreateUser(user models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserByUsername(username string) (models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}
