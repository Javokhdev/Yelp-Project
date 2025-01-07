package handlers

import (
	"net/http"
	"github.com/Javokhdev/Yelp-Project/internal/core/services"
	"github.com/Javokhdev/Yelp-Project/internal/domain"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// RegisterUser godoc
// @Summary      Register a new user
// @Description  Register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      domain.User  true  "User"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Router       /auth/register [post]
func (handler *AuthHandler) RegisterUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := handler.authService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}

// GetUserByID godoc
// @Summary      Get a user by ID
// @Description  Get a user by ID
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        userID  path      string  true  "User ID"
// @Success      200  {object}  domain.User
// @Failure      400  {object}  string
// @Router       /auth/{userID} [get]
func (handler *AuthHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("userID")
	user, err := handler.authService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Get all users
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  []domain.User
// @Failure      400  {object}  string
// @Router       /auth [get]
func (handler *AuthHandler) GetAllUsers(c *gin.Context) {
	users, err := handler.authService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Update a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      domain.User  true  "User"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      500  {object}  string
// @Router       /auth [put]
func (h *AuthHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.authService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete a user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        userID  path      string  true  "User ID"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// Failure      500  {object}  string
// @Router       /auth/{userID} [delete]
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("userID")
	if err := h.authService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
