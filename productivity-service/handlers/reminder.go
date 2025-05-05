package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nurcenter/productivity-service/client"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/models"
	"time"
)

func CreateReminder(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")
	var input struct {
		Title   string    `json:"title" binding:"required"`
		DueDate time.Time `json:"due_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	reminder := models.Reminder{
		UserID:  userID,
		Title:   input.Title,
		DueDate: input.DueDate,
	}

	if err := db.Create(&reminder).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create reminder"})
		return
	}

	c.JSON(201, reminder)
}
