package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nurcenter/productivity-service/client"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/models"
)

func CreateFinance(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")

	var input struct {
		Amount      float64 `json:"amount" binding:"required"`
		Description string  `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	finance := models.Finance{
		UserID:      userID,
		Amount:      input.Amount,
		Description: input.Description,
	}

	if err := db.Create(&finance).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create finance"})
		return
	}

	c.JSON(201, finance)
}
