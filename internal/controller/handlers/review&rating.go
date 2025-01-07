package handlers

import (
	"net/http"
	"github.com/Javokhdev/Yelp-Project/internal/core/services"
	"github.com/Javokhdev/Yelp-Project/internal/domain"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewService services.ReviewAndRatingService
}

func NewReviewHandler(reviewService *services.ReviewAndRatingService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: *reviewService,
	}
}

// Review godoc
// @Summary      Create a new review
// @Description  Create a new review
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        review  body  domain.Review  true  "Review details"
// @Success      200  {object}  domain.Review
// @Failure      400  {object}  string
// @Router       /review [post]
func (r *ReviewHandler) CreateReview(c *gin.Context) {
	var review domain.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reviewID, err := r.reviewService.CreateReview(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": reviewID})	
}

// Review godoc
// @Summary      Get all reviews
// @Description  Get all reviews
// @Tags         Review
// @Accept       json
// @Produce      json
// @Success      200  {object}  []domain.Review
// @Failure      400  {object}  string
// @Router       /review [get]
func (r *ReviewHandler) GetAllReviews(c *gin.Context) {
	reviews, err := r.reviewService.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// Review godoc
// @Summary      Get a review by ID
// @Description  Get a review by ID
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Review ID"
// @Success      200  {object}  domain.Review
// @Failure      400  {object}  string
// @Router       /review/{id} [get]
func (r *ReviewHandler) GetReviewByID(c *gin.Context) {
	reviewID := c.Param("id")
	review, err := r.reviewService.GetReviewByID(reviewID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

// Review godoc
// @Summary      Update a review
// @Description  Update a review
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        review  body  domain.Review  true  "Review details"
// @Success      200  {object}  domain.Review
// @Failure      400  {object}  string
// @Router       /review [put]
func (r *ReviewHandler) UpdateReview(c *gin.Context) {
	var review domain.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := r.reviewService.UpdateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

// Review godoc
// @Summary      Delete a review
// @Description  Delete a review
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Review ID"
// @Success      200  {object}  domain.Review
// @Failure      400  {object}  string
// @Router       /review/{id} [delete]
func (r *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewID := c.Param("id")
	if err := r.reviewService.DeleteReview(reviewID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
