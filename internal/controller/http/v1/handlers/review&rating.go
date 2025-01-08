package handlers

import (
	"net/http"
	"github.com/Javokhdev/Yelp-Project/internal/entity"

	"github.com/gin-gonic/gin"
)

// Review godoc
// @Summary      Create a new review
// @Description  Create a new review
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        review  body  entity.Review  true  "Review details"
// @Success      200  {object}  entity.Review
// @Failure      400  {object}  string
// @Router       /review [post]
func (h *Handler) CreateReview(c *gin.Context) {
	var review entity.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reviewID, err := h.UseCase.ReviewRepo.CreateReview(c, &review)
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
// @Success      200  {object}  []entity.Review
// @Failure      400  {object}  string
// @Router       /review [get]
func (h *Handler) GetAllReviews(c *gin.Context) {
	reviews, err := h.UseCase.ReviewRepo.GetAllReviews(c)
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
// @Success      200  {object}  entity.Review
// @Failure      400  {object}  string
// @Router       /review/{id} [get]
func (h *Handler) GetReviewByID(c *gin.Context) {
	reviewID := c.Param("id")
	review, err := h.UseCase.ReviewRepo.GetReviewByID(c, reviewID)
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
// @Param        review  body  entity.Review  true  "Review details"
// @Success      200  {object}  entity.Review
// @Failure      400  {object}  string
// @Router       /review [put]
func (h *Handler) UpdateReview(c *gin.Context) {
	var review entity.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UseCase.ReviewRepo.UpdateReview(c, &review); err != nil {
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
// @Success      200  {object}  entity.Review
// @Failure      400  {object}  string
// @Router       /review/{id} [delete]
func (h *Handler) DeleteReview(c *gin.Context) {
	reviewID := c.Param("id")
	if err := h.UseCase.ReviewRepo.DeleteReview(c, reviewID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

//Rating godoc
// @Summary      Get rating by business ID
// @Description  Get rating by business ID
// @Tags         Rating
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Business ID"
// @Success      200  {object}  entity.Rating
// @Failure      400  {object}  string
// @Router       /rating/{id} [get]
func (h *Handler) GetRatingByBusinessID(c *gin.Context) {
	businessID := c.Param("id")
	rating, err := h.UseCase.ReviewRepo.GetRatingByBusinessID(c, businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rating)
}
