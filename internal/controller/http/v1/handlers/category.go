package handlers

import (
	"net/http"
	"github.com/Javokhdev/Yelp-Project/internal/entity"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      entity.Category  true  "Category"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router       /categories [post]
func (h *Handler) CreateCategory(c *gin.Context) {
	var category entity.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryID, err := h.UseCase.CategoryRepo.CreateCategory(c, &category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category_id": categoryID})
}

// GetAllCategories godoc
// @Summary      Get all categories
// @Description  Get all categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  []entity.Category
// @Failure      400  {object}  string
// @Router       /categories [get]
func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := h.UseCase.CategoryRepo.GetAllCategories(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary      Get a category by ID
// @Description  Get a category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        categoryID  path      string  true  "Category ID"
// @Success      200  {object}  entity.Category
// @Failure      400  {object}  string
// @Router       /categories/{categoryID} [get]
func (h *Handler) GetCategoryByID(c *gin.Context) {
	var (
		req entity.Id
	)
	req.ID = c.Param("categoryID")

	category, err := h.UseCase.CategoryRepo.GetCategoryByID(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary      Update a category
// @Description  Update a category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      entity.Category  true  "Category"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /categories [put]	
func (h *Handler) UpdateCategory(c *gin.Context) {
	var category entity.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UseCase.CategoryRepo.UpdateCategory(c, &category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Delete a category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        categoryID  path      string  true  "Category ID"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /categories/{categoryID} [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
	var (
		req entity.Id
	)
	req.ID = c.Param("categoryID")
	if err := h.UseCase.CategoryRepo.DeleteCategory(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

