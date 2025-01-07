package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Javokhdev/Yelp-Project/config"
	"github.com/Javokhdev/Yelp-Project/internal/entity"
	"github.com/Javokhdev/Yelp-Project/pkg/logger"
	"github.com/Javokhdev/Yelp-Project/pkg/postgres"
	"github.com/google/uuid"
)

type ReviewRepo struct {
	pg     *postgres.Postgres
	config *config.Config
	logger *logger.Logger
}

// NewReviewRepo creates a new ReviewRepo instance.
func NewReviewRepo(pg *postgres.Postgres, config *config.Config, logger *logger.Logger) *ReviewRepo {
	return &ReviewRepo{
		pg:     pg,
		config: config,
		logger: logger,
	}
}

// CreateReview inserts a new review into the database.
func (r *ReviewRepo) CreateReview(ctx context.Context, review *entity.Review) (string, error) {
	review.ID = uuid.NewString()

	photosJSON, err := json.Marshal(review.Photos)
	if err != nil {
		return "", err
	}

	query, args, err := r.pg.Builder.Insert("reviews").
		Columns("id, business_id, user_id, rating, comment, photos").
		Values(review.ID, review.BusinessID, review.UserID, review.Rating, review.Comment, photosJSON).
		ToSql()
	if err != nil {
		return "", err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return review.ID, nil
}

// GetReviewByID retrieves a review by its ID.
func (r *ReviewRepo) GetReviewByID(ctx context.Context, reviewID string) (*entity.Review, error) {
	var (
		review     entity.Review
		photosJSON []byte
	)

	query, args, err := r.pg.Builder.
		Select("id, business_id, user_id, rating, comment, photos").
		From("reviews").
		Where("id = ?", reviewID).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.pg.Pool.QueryRow(ctx, query, args...).
		Scan(&review.ID, &review.BusinessID, &review.UserID, &review.Rating, &review.Comment, &photosJSON)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(photosJSON, &review.Photos); err != nil {
		return nil, err
	}

	return &review, nil
}

// GetAllReviewsByBusinessID retrieves all reviews for a specific business.
func (r *ReviewRepo) GetAllReviewsByBusinessID(ctx context.Context, businessID string) ([]*entity.Review, error) {
	query, args, err := r.pg.Builder.
		Select("id, business_id, user_id, rating, comment, photos").
		From("reviews").
		Where("business_id = ?", businessID).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*entity.Review
	for rows.Next() {
		var (
			review     entity.Review
			photosJSON []byte
		)

		err := rows.Scan(&review.ID, &review.BusinessID, &review.UserID, &review.Rating, &review.Comment, &photosJSON)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(photosJSON, &review.Photos); err != nil {
			return nil, err
		}

		reviews = append(reviews, &review)
	}

	return reviews, nil
}


// GetAllReviews retrieves all reviews from the database.
func (r *ReviewRepo) GetAllReviews(ctx context.Context) ([]*entity.Review, error) {
	query, args, err := r.pg.Builder.
		Select("id, business_id, user_id, rating, comment, photos").
		From("reviews").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*entity.Review
	for rows.Next() {
		var (
			review     entity.Review
			photosJSON []byte
		)

		err := rows.Scan(&review.ID, &review.BusinessID, &review.UserID, &review.Rating, &review.Comment, &photosJSON)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(photosJSON, &review.Photos); err != nil {
			return nil, err
		}

		reviews = append(reviews, &review)
	}

	return reviews, nil
}

// UpdateReview updates an existing review.
func (r *ReviewRepo) UpdateReview(ctx context.Context, review *entity.Review) error {
	photosJSON, err := json.Marshal(review.Photos)
	if err != nil {
		return err
	}

	query, args, err := r.pg.Builder.
		Update("reviews").
		SetMap(map[string]interface{}{
			"rating":   review.Rating,
			"comment":  review.Comment,
			"photos":   photosJSON,
			"updated_at": time.Now(),
		}).
		Where("id = ?", review.ID).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	return err
}

// DeleteReview deletes a review by its ID.
func (r *ReviewRepo) DeleteReview(ctx context.Context, reviewID string) error {
	query, args, err := r.pg.Builder.Delete("reviews").Where("id = ?", reviewID).ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	return err
}

// GetRatingByBusinessID retrieves the average rating and total ratings for a specific business.
func (r *ReviewRepo) GetRatingByBusinessID(ctx context.Context, businessID string) (*entity.Rating, error) {
	query, args, err := r.pg.Builder.
		Select("business_id, AVG(rating) AS avg_rating, COUNT(rating) AS total_ratings").
		From("reviews").
		Where("business_id = ?", businessID).
		GroupBy("business_id").
		ToSql()
	if err != nil {
		return nil, err
	}

	var rating entity.Rating
	err = r.pg.Pool.QueryRow(ctx, query, args...).
		Scan(&rating.BusinessID, &rating.AvgRating, &rating.TotalRatings)
	if err != nil {
		return nil, err
	}

	return &rating, nil
}



