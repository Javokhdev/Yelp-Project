package repositories

import (
	"database/sql"
	domain "github.com/Javokhdev/Yelp-Project/internal/domain"

	"github.com/google/uuid"
)

type ReviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	return &ReviewRepo{
		db: db,
	}
}

func (r *ReviewRepo) CreateReview(review *domain.Review) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO reviews (id, business_id, user_id, rating, comment, photos)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	var reviewID string
	err := r.db.QueryRow(query, id, review.BusinessID, review.UserID, review.Rating, review.Comment, review.Photos).Scan(&reviewID)
	if err != nil {
		return "", err
	}
	return reviewID, nil
}

func (r *ReviewRepo) GetReviewByID(reviewID string) (*domain.Review, error) {
	var review domain.Review

	query := `SELECT id, business_id, user_id, rating, comment, photos FROM reviews WHERE id = $1`
	err := r.db.QueryRow(query, reviewID).Scan(&review.ID, &review.BusinessID, &review.UserID, &review.Rating, &review.Comment, &review.Photos)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepo) GetAllReviews() ([]*domain.Review, error) {
	query := `SELECT id, business_id, user_id, rating, comment, photos FROM reviews`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*domain.Review
	for rows.Next() {
		review := &domain.Review{}
		err := rows.Scan(&review.ID, &review.BusinessID, &review.UserID, &review.Rating, &review.Comment, &review.Photos)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil	
}

func (r *ReviewRepo) UpdateReview(review *domain.Review) error {
	_, err := r.db.Exec(`
		UPDATE reviews
		SET rating = $1, comment = $2, photos = $3
		WHERE id = $4
	`, review.Rating, review.Comment, review.Photos, review.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewRepo) DeleteReview(reviewID string) error {
	_, err := r.db.Exec("DELETE FROM reviews WHERE id = $1", reviewID)
	if err != nil {
		return err
	}
	return nil
}