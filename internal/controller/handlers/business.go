package handlers

import (
	"net/http"
	"yalp/internal/core/services"
	"yalp/internal/domain"

	"github.com/gin-gonic/gin"
)

type BusinessHandler struct {
	businessService *services.BusinessService
}

func NewBusinessHandler(businessService *services.BusinessService) *BusinessHandler {
	return &BusinessHandler{
		businessService: businessService,
	}
}

// Business godoc
// @Summary      Create a new business
// @Description  Create a new business
// @Tags         Business
// @Accept       json
// @Produce      json
// @Param        business  body  domain.Business  true  "Business details"
// @Success      200  {object}  string "Succefully"
// @Failure      400  {object}  string
// Failure      500  {object}  string
// @Router       /business [post]
func (handler *BusinessHandler) CreateBusiness(c *gin.Context) {
	var business domain.Business
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	businessID, err := handler.businessService.CreateBusiness(&business)
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
func (handler *BusinessHandler) GetBusinessByID(c *gin.Context) {
	businessID := c.Param("businessID")
	business, err := handler.businessService.GetBusinessByID(businessID)
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
func (handler *BusinessHandler) GetAllBusinesses(c *gin.Context) {
	businesses, err := handler.businessService.GetAllBusinesses()
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
// @Param        business  body  domain.Business  true  "Business details"
// @Success      200  {object}  string "Succefully"
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /business [put]
func (handler *BusinessHandler) UpdateBusiness(c *gin.Context) {
	var business domain.Business
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := handler.businessService.UpdateBusiness(&business); err != nil {
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
func (handler *BusinessHandler) DeleteBusiness(c *gin.Context) {
	businessID := c.Param("businessID")
	if err := handler.businessService.DeleteBusiness(businessID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Business deleted successfully"})
}	