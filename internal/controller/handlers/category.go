package handlers

import (
	"net/http"
	"yalp/internal/core/services"
	"yalp/internal/domain"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: *categoryService,
	}
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      domain.Category  true  "Category"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router       /categories [post]
func (handler *CategoryHandler) CreateCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryID, err := handler.categoryService.CreateCategory(&category)
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
// @Success      200  {object}  []domain.Category
// @Failure      400  {object}  string
// @Router       /categories [get]
func (handler *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := handler.categoryService.GetAllCategories()
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
// @Success      200  {object}  domain.Category
// @Failure      400  {object}  string
// @Router       /categories/{categoryID} [get]
func (handler *CategoryHandler) GetCategoryByID(c *gin.Context) {
	categoryID := c.Param("categoryID")
	category, err := handler.categoryService.GetCategoryByID(categoryID)
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
// @Param        category  body      domain.Category  true  "Category"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /categories [put]	
func (handler *CategoryHandler) UpdateCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := handler.categoryService.UpdateCategory(&category); err != nil {
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
func (handler *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("categoryID")
	if err := handler.categoryService.DeleteCategory(categoryID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

