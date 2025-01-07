package services

import (
	"github.com/Javokhdev/Yelp-Project/internal/core/repositories"
	"github.com/Javokhdev/Yelp-Project/internal/domain"
)

type ReviewAndRatingService struct {
	repo repositories.ReviewRepo
}

func NewReviewAndRatingService(repo repositories.ReviewRepo) *ReviewAndRatingService {
	return &ReviewAndRatingService{
		repo: repo,
	}
}

func (s *ReviewAndRatingService) CreateReview(review *domain.Review) (string, error) {
	return s.repo.CreateReview(review)
}

func (s *ReviewAndRatingService) GetReviewByID(reviewID string) (*domain.Review, error) {
	return s.repo.GetReviewByID(reviewID)
}

func (s *ReviewAndRatingService) GetAllReviews() ([]*domain.Review, error) {
	return s.repo.GetAllReviews()
}	

func (s *ReviewAndRatingService) UpdateReview(review *domain.Review) error {
	return s.repo.UpdateReview(review)
}

func (s *ReviewAndRatingService) DeleteReview(reviewID string) error {
	return s.repo.DeleteReview(reviewID)
}