package services

import (
	"yalp/internal/core/repositories"
	"yalp/internal/domain"
)

type BusinessService struct {
	repo repositories.BusinessRepo
}

func NewBusinessService(repo repositories.BusinessRepo) *BusinessService {
	return &BusinessService{
		repo: repo,
	}
}

func (s *BusinessService) CreateBusiness(business *domain.Business) (string, error) {
	return s.repo.CreateBusiness(business)
}

func (s *BusinessService) GetBusinessByID(businessID string) (*domain.Business, error) {
	return s.repo.GetBusinessByID(businessID)
}

func (s *BusinessService) GetAllBusinesses() ([]*domain.Business, error) {
	return s.repo.GetAllBusinesses()
}

func (s *BusinessService) UpdateBusiness(business *domain.Business) error {
	return s.repo.UpdateBusiness(business)
}

func (s *BusinessService) DeleteBusiness(businessID string) error {
	return s.repo.DeleteBusiness(businessID)
}