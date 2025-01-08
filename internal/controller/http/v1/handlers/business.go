package handlers

import (
	"net/http"

	"github.com/Javokhdev/Yelp-Project/internal/entity"
	"github.com/gin-gonic/gin"
)

// Business godoc
// @Summary      Create a new business
// @Description  Create a new business
// @Tags         Business
// @Accept       json
// @Produce      json
// @Param        business  body  entity.Business  true  "Business details"
// @Success      200  {object}  string "Succefully"
// @Failure      400  {object}  string
// Failure      500  {object}  string
// @Router       /business [post]
func (h *Handler) CreateBusiness(c *gin.Context) {
	var business entity.Business
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	businessID, err := h.UseCase.BusinessRepo.CreateBusiness(c, &business)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"business_id": businessID})
}

// Business godoc
// @Summary      Get a business by ID
// @Description  Get a business by ID
// @Tags         Business
// @Accept       json
// @Produce      json
// @Param        businessID  path      string  true  "Business ID"
// @Success      200  {object}  string "Succefully"
// @Failure      400  {object}  string
// Failure      500  {object}  string
// @Router       /business/{businessID} [get]
func (h *Handler) GetBusinessByID(c *gin.Context) {
	businessID := c.Param("businessID")
	business, err := h.UseCase.BusinessRepo.GetBusinessByID(c,businessID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, business)
}

// Business godoc
// @Summary      Get all businesses
// @Description  Get all businesses
// @Tags         Business
// @Accept       json
// @Produce      json
// @Success      200  {object}  string "Succefully"
// @Failure      400  {object}  string
// @Router       /business [get]
func (h *Handler) GetAllBusinesses(c *gin.Context) {
	businesses, err := h.UseCase.BusinessRepo.GetAllBusinesses(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, businesses)
}

// Business godoc
// @Summary      Update a business
// @Description  Update a business
// @Tags         Business
// @Accept       json
// @Produce      json
// @Param        business  body  entity.Business  true  "Business details"
// @Success      200  {object}  string "Succefully"
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /business [put]
func (h *Handler) UpdateBusiness(c *gin.Context) {
	var business entity.Business
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UseCase.BusinessRepo.UpdateBusiness(c, &business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Business updated successfully"})
}

// Business godoc
// @Summary      Delete a business
// @Description  Delete a business
// @Tags         Business
// @Accept       json
// @Produce      json
// @Param        businessID  path      string  true  "Business ID"
// @Success      200  {object}  string "Succefully"
// @Failure      400  {object}  string
// @Router       /business/{businessID} [delete]
func (h *Handler) DeleteBusiness(c *gin.Context) {
	businessID := c.Param("businessID")
	if err := h.UseCase.BusinessRepo.DeleteBusiness(c, businessID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Business deleted successfully"})
}
