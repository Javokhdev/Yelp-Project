package services

import (
	"github.com/Javokhdev/Yelp-Project/internal/domain"
	"github.com/Javokhdev/Yelp-Project/internal/core/repositories"
)

type AuthService struct {
	repo repositories.UserRepo
}

func NewAuthService(repo repositories.UserRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) RegisterUser(user *domain.User) (string, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByID(userID string) (*domain.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *AuthService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *AuthService) UpdateUser(user *domain.User) error {
	return s.repo.UpdateUser(user)
}

func (s *AuthService) DeleteUser(userID string) error {
	return s.repo.DeleteUser(userID)
}