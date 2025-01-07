package storage

import "github.com/Javokhdev/Yelp-Project/internal/domain"

type BusinessInterface interface {
	CreateBusiness(business *domain.Business) (string, error)
	GetBusinessByID(businessID string) (*domain.Business, error)
	GetAllBusinesses() ([]*domain.Business, error)
	UpdateBusiness(business *domain.Business) error
	DeleteBusiness(businessID string) error
}

type ReviewInterface interface {
	CreateReview(review *domain.Review) (string, error)
	GetReviewByID(reviewID string) (*domain.Review, error)
	GetAllReviewsByBusinessID(businessID string) ([]*domain.Review, error)
	UpdateReview(review *domain.Review) error
	DeleteReview(reviewID string) error
}

type RatingInterface interface {
	GetRatingByBusinessID(businessID string) (*domain.Rating, error)
}

type AuthInterface interface {
	CreatePasswordResetRequest(request *domain.PasswordResetRequest) error
	GetPasswordResetRequestByID(requestID string) (*domain.PasswordResetRequest, error)
	UpdatePasswordResetRequest(request *domain.PasswordResetRequest) error
	DeletePasswordResetRequest(requestID string) error
}

type UserInterface interface {
	CreateUser(user *domain.User) (string, error)
	GetUserByID(userID string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(userID string) error
}