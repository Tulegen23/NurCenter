package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nurcenter/productivity-service/client"
	"nurcenter/productivity-service/config"
	"nurcenter/productivity-service/models"
	"strconv"
	"time"
)

func GetGoals(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	var goals []models.Goal
	if err := db.Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&goals).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to get goals"})
		return
	}

	var total int64
	db.Model(&models.Goal{}).Where("user_id = ?", userID).Count(&total)

	c.JSON(200, gin.H{
		"goals": goals,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func CreateGoal(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	userID := c.GetUint("user_id")

	var input struct {
		Title    string    `json:"title" binding:"required"`
		Deadline time.Time `json:"deadline"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := client.VerifyUser(userID, cfg); err != nil {
		c.JSON(500, gin.H{"error": "Failed to verify user"})
		return
	}

	goal := models.Goal{
		UserID:   userID,
		Title:    input.Title,
		Deadline: input.Deadline,
	}

	if err := db.Create(&goal).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create goal"})
		return
	}

	c.JSON(201, goal)
}
