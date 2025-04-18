package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucamienert/flashcards/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// GetMe godoc
// @Summary Get current authenticated user
// @Description Get the details of the currently authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} gin.H{"error": "message"}
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Security BearerAuth
// @Router /user/me [get]
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
