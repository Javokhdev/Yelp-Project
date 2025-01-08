package usecase

import (
	"github.com/Javokhdev/Yelp-Project/internal/entity"
	"context"
)

type (
	
 	BusinessI interface {
		CreateBusiness(ctx context.Context, business *entity.Business) (string, error)
		GetBusinessByID(ctx context.Context, businessID string) (*entity.Business, error)
		GetAllBusinesses(ctx context.Context) ([]*entity.Business, error)
		UpdateBusiness(ctx context.Context,business *entity.Business) error
		DeleteBusiness(ctx context.Context, businessID string) error
	}	

 	ReviewI interface {
		CreateReview(ctx context.Context, review *entity.Review) (string, error)
		GetReviewByID(ctx context.Context, reviewID string) (*entity.Review, error)
		GetAllReviewsByBusinessID(ctx context.Context, businessID string) ([]*entity.Review, error)
		GetAllReviews(ctx context.Context) ([]*entity.Review, error)
		UpdateReview(ctx context.Context, review *entity.Review) error
		DeleteReview(ctx context.Context, reviewID string) error
		GetRatingByBusinessID(ctx context.Context, businessID string) (*entity.Rating, error)
	}	


	// UserRepo -.
	UserRepoI interface {
		Create(ctx context.Context, req entity.User) (entity.User, error)
		GetSingle(ctx context.Context, req entity.UserSingleRequest) (entity.User, error)
		GetList(ctx context.Context, req entity.GetListFilter) (entity.UserList, error)
		Update(ctx context.Context, req entity.User) (entity.User, error)
		Delete(ctx context.Context, req entity.Id) error
		UpdateField(ctx context.Context, req entity.UpdateFieldRequest) (entity.RowsEffected, error)
	}

	// SessionRepo -.
	SessionRepoI interface {
		Create(ctx context.Context, req entity.Session) (entity.Session, error)
		GetSingle(ctx context.Context, req entity.Id) (entity.Session, error)
		GetList(ctx context.Context, req entity.GetListFilter) (entity.SessionList, error)
		Update(ctx context.Context, req entity.Session) (entity.Session, error)
		Delete(ctx context.Context, req entity.Id) error
		UpdateField(ctx context.Context, req entity.UpdateFieldRequest) (entity.RowsEffected, error)
	}

	CategoryRepoI interface {
		CreateCategory(ctx context.Context, req *entity.Category) (string, error)
		GetAllCategories(ctx context.Context) ([]*entity.Category, error)
		GetCategoryByID(ctx context.Context, req entity.Id) (*entity.Category, error)
		UpdateCategory(ctx context.Context, req *entity.Category) (error)
		DeleteCategory(ctx context.Context, req entity.Id) error
	}
)