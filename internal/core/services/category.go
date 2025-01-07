package services

import (
	"github.com/Javokhdev/Yelp-Project/internal/core/repositories"
	"github.com/Javokhdev/Yelp-Project/internal/domain"
)

type CategoryService struct {
	repo repositories.CategoryRepo
}

func NewCategoryService(repo repositories.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *domain.Category) (string, error) {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetCategoryByID(categoryID string) (*domain.Category, error) {
	return s.repo.GetCategoryByID(categoryID)
}

func (s *CategoryService) GetAllCategories() ([]*domain.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) UpdateCategory(category *domain.Category) error {
	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(categoryID string) error {
	return s.repo.DeleteCategory(categoryID)
}